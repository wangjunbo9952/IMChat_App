const currentUserId = document.getElementById('currentUserId').value;
const currentUserAccount = document.getElementById('currentUserAccount').value;


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

    if (message) {
        console.log('发送消息:', message);
        input.value = '';
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

            // 使用 map 遍历好友列表，为每个好友创建 HTML
            const friendsHTML = result.friends.map(friend => `
    <div class="friend-item" onclick="showChatHistory(${friend.ID}, '${friend.username}')">
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

async function showChatHistory(friendId, friendName) {
    // 更新聊天头部显示的好友名称
    document.querySelector('.chat-header h2').textContent = friendName;

    try {
        const msgReq = {
            messageType: 1,  // 假设 1 代表获取聊天记录
            userId: Number(currentUserId),
            targetID: Number(friendId)
        };

        const response = await fetch('http://127.0.0.1:9090/user/history', {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(msgReq)
        });

        const result = await response.json();
        const chatMessages = document.querySelector('.chat-messages');

        if (result.success && result.messages) {
            const messagesHTML = result.messages.map(msg => {
                const numCurrentUserId = Number(currentUserId);
                const numSenderId = Number(msg.senderId);

                console.log('消息发送者ID:', numSenderId);
                console.log('当前用户ID:', numCurrentUserId);
                console.log('是否是发送消息:', numSenderId === numCurrentUserId);

                return `
                <div class="message ${numSenderId === numCurrentUserId ? 'sent' : 'received'}">
                    <div class="message-content">${msg.content}</div>
                    <div class="message-time">${new Date(msg.timestamp).toLocaleString()}</div>
                </div>
            `;
            }).join('');

            chatMessages.innerHTML = messagesHTML;
            chatMessages.scrollTop = chatMessages.scrollHeight;
        } else {
            chatMessages.innerHTML = '<div class="no-messages">暂无聊天记录</div>';
        }
    } catch (error) {
        console.error('获取聊天记录失败:', error);
    }
}

let ws = null;

function connectWebSocket() {
    const user = currentUserAccount;
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

        ws.onopen = function() {
            console.log('WebSocket连接已建立!');
            wsIcon.className = 'nav-icon connected';
        };

        ws.onmessage = function(event) {
            console.log('收到消息:', event.data);
            const reader = new FileReader();
            reader.onloadend = function () {
                console.log(reader.result);  // 输出：Blob 数据
            };

            reader.readAsText(event.data);  // 将 Blob 内容作为文本读取

        };

        ws.onerror = function(error) {
            console.error('WebSocket错误:', error);
            wsIcon.className = 'nav-icon disconnected';
        };

        ws.onclose = function() {
            console.log('WebSocket连接已关闭');
            wsIcon.className = 'nav-icon disconnected';
        };

        window.chatWebSocket = ws;

    } catch (error) {
        console.error('创建WebSocket连接失败:', error);
        wsIcon.className = 'nav-icon disconnected';
    }
}