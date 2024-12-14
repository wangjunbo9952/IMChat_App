const currentUserId = document.getElementById('currentUserId').value;
const currentUserAccount = document.getElementById('currentUserAccount').value;
let ws = null;
let currentChatUser = null; // 当前聊天对象


// 首先定义Message结构
const messageProto = `
syntax = "proto3";

message Message {
    string avatar = 1;       
    string fromUsername = 2; 
    string from = 3;         
    string to = 4;           
    string content = 5;      
    int32 contentType = 6;   
    string type = 7;         
    int32 messageType = 8;   
    string url = 9;          
    string fileSuffix = 10;  
    bytes file = 11;        
}`;

// 在页面加载时初始化protobuf
let Message;
try {
    const root = protobuf.parse(messageProto).root;
    Message = root.lookupType("Message");
    console.log("Protobuf 初始化成功");
} catch (err) {
    console.error("Protobuf 初始化失败:", err);
}

function toggleDropdown() {
    const dropdown = document.getElementById('dropdownMenu');
    dropdown.classList.toggle('show');
}

// 点击其他地方关闭下拉菜单
document.addEventListener('click', function(event) {
    const dropdown = document.getElementById('dropdownMenu');
    const addButton = event.target.closest('.add-button');

    if (!addButton && dropdown.classList.contains('show')) {
        dropdown.classList.remove('show');
    }
});

// 导航图标切换
document.querySelectorAll('.nav-icon').forEach(icon => {
    icon.addEventListener('click', function() {
        document.querySelectorAll('.nav-icon').forEach(i => i.classList.remove('active'));
        this.classList.add('active');
    });
});

// 发送消息功能
function sendMessage() {
    const input = document.getElementById('messageInput');
    const message = input.value.trim();

    if (!message || !currentChatUser || !ws) {
        console.error('无法发送消息：消息为空或未选择聊天对象或WebSocket未连接');
        return;
    }

    try {
        const msgData = {
            avatar: '/static/default-avatar.jpg',
            fromUsername: currentUserAccount,
            from: currentUserId,
            to: currentChatUser,
            content: message,
            contentType: 1,
            type: 'chat',
            messageType: 1,
            url: '',
            fileSuffix: '',
            file: new Uint8Array(),
            timestamp: Date.now() // 添加时间戳
        };

        // 验证消息格式
        const errMsg = Message.verify(msgData);
        if (errMsg) {
            throw Error(errMsg);
        }

        // 创建消息实例
        const protoMessage = Message.create(msgData);

        // 序列化消息
        const buffer = Message.encode(protoMessage).finish();

        // 发送二进制数据
        ws.send(buffer);

        // 显示发送的消息
        appendMessage(msgData, true);

        // 清空输入框
        input.value = '';

    } catch (error) {
        console.error('发送消息失败:', error);
    }
}

// 显示消息函数
function appendMessage(msg, isSent) {
    const chatMessages = document.querySelector('.chat-messages');
    const messageDiv = document.createElement('div');
    messageDiv.className = `message ${isSent ? 'sent' : 'received'}`;

    // 处理时间显示
    let timeStr;
    if (msg.timestamp) {
        const date = new Date(Number(msg.timestamp));
        timeStr = formatMessageTime(date);
    } else {
        timeStr = formatMessageTime(new Date());
    }

    messageDiv.innerHTML = `
        <div class="message-content">${msg.content}</div>
        <div class="message-time">${timeStr}</div>
    `;

    chatMessages.appendChild(messageDiv);
    chatMessages.scrollTop = chatMessages.scrollHeight;
}

// 添加时间格式化函数
function formatMessageTime(date) {
    try {
        const now = new Date();
        const today = new Date(now.getFullYear(), now.getMonth(), now.getDate());
        const yesterday = new Date(today);
        yesterday.setDate(yesterday.getDate() - 1);

        const hours = date.getHours().toString().padStart(2, '0');
        const minutes = date.getMinutes().toString().padStart(2, '0');
        const timeStr = `${hours}:${minutes}`;

        if (date >= today) {
            return timeStr; // 今天的消息只显示时间
        } else if (date >= yesterday) {
            return `昨天 ${timeStr}`; // 昨天的消息
        } else {
            // 更早的消息显示完整日期
            const month = (date.getMonth() + 1).toString().padStart(2, '0');
            const day = date.getDate().toString().padStart(2, '0');
            return `${month}-${day} ${timeStr}`;
        }
    } catch (error) {
        console.error('时间格式化失败:', error);
        return '时间未知';
    }
}

// 处理收到的消息
function handleReceivedMessage(event) {
    try {
        // 直接使用ArrayBuffer数据
        const buffer = new Uint8Array(event.data);
        const msg = Message.decode(buffer);
        console.log('收到消息:', msg);

        // 显示消息
        appendMessage({
            content: msg.content,
            timestamp: new Date().getTime(),
            from: msg.from,
            to: msg.to
        }, msg.from === currentUserId);

    } catch (error) {
        console.error('处理消息失败:', error);
    }
}

// 回车发送功能
document.getElementById('messageInput').addEventListener('keypress', function(e) {
    if (e.key === 'Enter') {
        sendMessage();
    }
});

// 获取模态框元素
const modal = document.getElementById('addFriendModal');
const closeBtn = document.querySelector('.close-btn');

// 点击添加好友时打开模态框
document.querySelector('.dropdown-item:first-child').addEventListener('click', function(e) {
    e.preventDefault();
    modal.style.display = 'block';
    document.getElementById('friendAccount').focus();
});

// 点击关闭按钮关闭模态框
closeBtn.addEventListener('click', function() {
    modal.style.display = 'none';
});

// 点击模态框外部关闭
window.addEventListener('click', function(e) {
    if (e.target == modal) {
        modal.style.display = 'none';
    }
});

// 搜索好友功能
// 在 chat.html 的 script 标签中修改 searchFriend 函数
async function searchFriend() {
    const account = document.getElementById('friendAccount').value.trim();
    const searchResult = document.getElementById('searchResult');

    if (!account) {
        alert('请输入好友账号');
        return;
    }

    try {
        searchResult.innerHTML = '正在搜索...';

        const response = await fetch(`http://127.0.0.1:9090/user/searchuser?account=${account}`, {
            method: 'GET',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        const result = await response.json();

        console.log('搜索到的用户ID:', result.id);

        if (result.success) {
            // 显示搜索结果
            searchResult.innerHTML = `
                <div class="user-card">
                    <img src="${result.avatar || '/static/default-avatar.jpg'}" alt="头像" class="user-avatar">
                    <div class="user-info">
                        <h4>${result.username}</h4>
                        <p>${result.account}</p>
                    </div>
                    <button class="add-friend-btn" onclick="addFriend(${result.id})">
                        <i class="fas fa-user-plus"></i>
                        添加
                    </button>
                </div>
            `;
        } else {
            searchResult.innerHTML = '<p class="no-result">未找到该用户</p>';
        }
    } catch (error) {
        console.error('搜索失败:', error);
        searchResult.innerHTML = '<p class="error-msg">搜索失败，请重试</p>';
    }
}

// 添加好友函数
async function addFriend(targetUserId) {

    const numUserId = Number(currentUserId);
    const numFriendId = Number(targetUserId);

    try {
        const response = await fetch('http://127.0.0.1:9090/user/adduser', {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                userId: numUserId,
                friendId: numFriendId
            })
        });

        const result = await response.json();

        if (result.success) {
            alert('好友请求已发送');
            document.getElementById('addFriendModal').style.display = 'none';
        } else {
            alert(result.message || '添加好友失败');
        }
    } catch (error) {
        console.error('添加好友失败:', error);
        alert('添加好友失败，请重试');
    }
}

// 添加回车搜索功能
document.getElementById('friendAccount').addEventListener('keypress', function(e) {
    if (e.key === 'Enter') {
        searchFriend();
    }
});

async function loadFriendsList() {
    try {
        console.log('加载好友列表，当前用户ID:', currentUserId);

        const response = await fetch(`http://127.0.0.1:9090/user/getfriends?userId=${currentUserId}`, {
            method: 'GET',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        const result = await response.json();
        console.log('服务器返回的好友列表:', result);

        if (result.success) {
            const container = document.getElementById('friendsContainer');

            // 检查是否有好友
            if (!result.friends || result.friends.length === 0) {
                container.innerHTML = '<div class="no-friends">暂无好友</div>';
                return;
            }

            // 修改这里，为每个好友添加点击事件
            const friendsHTML = result.friends.map(friend => `
                <div class="friend-item" onclick="selectChatUser('${friend.ID}', '${friend.username}')">
                    <img src="${friend.avatar || '/static/default-avatar.jpg'}" alt="头像" class="friend-avatar">
                    <div class="friend-info">
                        <div class="friend-name">${friend.username}</div>
                        <div class="friend-account">${friend.account}</div>
                    </div>
                </div>
            `).join('');

            // 将生成的 HTML 插入到容器中
            container.innerHTML = friendsHTML;
        }
    } catch (error) {
        console.error('请求失败:', error);
    }
}

// chat.js

async function loadChatHistory(friendId) {
    try {
        // 显示加载中的提示
        const chatMessages = document.querySelector('.chat-messages');
        chatMessages.innerHTML = '<div class="loading-messages">加载聊天记录中...</div>';

        const msgReq = {
            messageType: 1,  // 假设 1 代表获取聊天记录
            userId: Number(currentUserId),
            targetID: Number(friendId)
        };

        console.log('请求历史消息，参数:', msgReq);

        const response = await fetch('http://127.0.0.1:9090/user/history', {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(msgReq)
        });

        const result = await response.json();
        console.log('获取到的历史消息:', result);

        if (result.success && result.messages) {
            // 清空加载提示
            chatMessages.innerHTML = '';

            // 如果没有消息记录
            if (result.messages.length === 0) {
                chatMessages.innerHTML = '<div class="no-messages">暂无聊天记录</div>';
                return;
            }

            // 对消息按时间戳进行排序
            const sortedMessages = result.messages.sort((a, b) => {
                // 将字符串时间转换为时间戳进行比较
                const timeA = new Date(a.timestamp).getTime();
                const timeB = new Date(b.timestamp).getTime();
                return timeA - timeB; // 升序排列，最早的消息在前
            });

            console.log('排序后的消息:', sortedMessages);

            // 显示排序后的消息
            sortedMessages.forEach(msg => {
                const isSent = Number(msg.senderId) === Number(currentUserId);
                appendMessage({
                    content: msg.content,
                    timestamp: msg.timestamp,
                    from: msg.senderId,
                    to: msg.receiverId
                }, isSent);
            });

            // 滚动到最新消息
            chatMessages.scrollTop = chatMessages.scrollHeight;
        } else {
            chatMessages.innerHTML = '<div class="error-message">加载聊天记录失败</div>';
        }
    } catch (error) {
        console.error('加载聊天记录失败:', error);
        document.querySelector('.chat-messages').innerHTML =
            '<div class="error-message">加载聊天记录失败，请重试</div>';
    }
}


function connectWebSocket() {
    const user = currentUserId;
    if (!user) {
        alert('未找到用户信息');
        return;
    }

    if (ws && ws.readyState === WebSocket.OPEN) {
        alert('WebSocket已连接');
        return;
    }

    const wsIcon = document.getElementById('ws-connect');
    wsIcon.className = 'nav-icon connecting';

    const wsUrl = `ws://127.0.0.1:9090/ws?user=${encodeURIComponent(user)}`;
    console.log('尝试连接WebSocket:', wsUrl);

    try {
        ws = new WebSocket(wsUrl);

        ws.binaryType = 'arraybuffer';

        ws.onopen = function() {
            console.log('WebSocket连接已建立!');
            const wsIcon = document.getElementById('ws-connect');
            wsIcon.className = 'nav-icon connected';
        };

        ws.onmessage = handleReceivedMessage;

        ws.onerror = function(error) {
            console.error('WebSocket错误:', error);
            const wsIcon = document.getElementById('ws-connect');
            wsIcon.className = 'nav-icon disconnected';
        };

        ws.onclose = function() {
            console.log('WebSocket连接已关闭');
            const wsIcon = document.getElementById('ws-connect');
            wsIcon.className = 'nav-icon disconnected';
        };

        window.chatWebSocket = ws;

    } catch (error) {
        console.error('创建WebSocket连接失败:', error);
        wsIcon.className = 'nav-icon disconnected';
    }
}

// 选择聊天对象时的处理
function selectChatUser(userId, userName) {
    // 保存当前聊天对象ID
    currentChatUser = userId;

    // 更新聊天窗口标题
    document.querySelector('.chat-header h2').textContent = userName;

    // 清空聊天记录区域
    const chatMessages = document.querySelector('.chat-messages');
    chatMessages.innerHTML = '';

    // 激活输入框
    const messageInput = document.getElementById('messageInput');
    messageInput.disabled = false;
    messageInput.placeholder = `发送消息给 ${userName}...`;
    messageInput.focus();

    // 高亮选中的好友
    document.querySelectorAll('.friend-item').forEach(item => {
        item.classList.remove('active');
    });
    event.currentTarget.classList.add('active');

    // 加载历史消息
    loadChatHistory(userId);
}
