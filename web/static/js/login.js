function switchTab(tab) {
    document.querySelectorAll('.tab').forEach(t => t.classList.remove('active'));
    document.querySelectorAll('.form').forEach(f => f.classList.remove('active'));

    document.querySelector(`.tab:${tab === 'login' ? 'first-child' : 'last-child'}`).classList.add('active');
    document.getElementById(tab).classList.add('active');
}
// 表单提交处理函数
async function handleSubmit(formType) {
    try {
        const form = document.getElementById(formType);
        const inputs = form.getElementsByTagName('input');
        const formData = {};

        // 表单验证（放在数据收集之前）
        if (formType === 'register') {
            const password = form.querySelector('input[name="password"]').value.trim();
            const confirmPassword = form.querySelector('input[name="confirmPassword"]').value.trim();

            if (password !== confirmPassword) {
                alert('两次输入的密码不一致！');
                return;
            }
        }

        // 收集表单数据
        for (let input of inputs) {
            // 跳过确认密码字段
            if (input.name !== 'confirmPassword') {
                formData[input.name] = input.value;
            }
        }

        console.log('发送到后端的数据：', formData);

        // 定义 API 端点
        const apiEndpoints = {
            login: 'http://127.0.0.1:9090/user/login',
            register: 'http://127.0.0.1:9090/user/register'
        };

        // 发送请求到后端
        const response = await fetch(apiEndpoints[formType], {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        });

        if (!response.ok) {
            throw new Error('请求失败');
        }

        const result = await response.json();

        // 处理响应
        if (result.success) {
            if (formType === 'login') {
                let baseUrl = 'http://127.0.0.1:9090/chat/';
                let href = baseUrl + encodeURIComponent(result.account);  // 预先生成完整的 URL

                console.log("Generated URL:", href);

                window.location.href = href;
            } else {
                // 注册成功，切换到登录表单
                alert('注册成功！');
                switchTab('login');
                // 可选：自动填充用户名
                document.querySelector('#login input[name="account"]').value = formData.account;
            }
        } else {
            alert(result.message || '操作失败，请重试！');
        }

    } catch (error) {
        console.error('Error:', error);
        alert('操作失败，请检查网络连接！');
    }
}

// 添加表单回车提交功能
document.querySelectorAll('.form').forEach(form => {
    form.addEventListener('keypress', (e) => {
        if (e.key === 'Enter') {
            const formType = form.id;
            handleSubmit(formType);
        }
    });
});