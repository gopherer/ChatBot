document.getElementById('login-button').addEventListener('click', function() {
    // 获取用户输入
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const mode = document.getElementById('mode').value;

    // 创建要发送的数据对象
    const data = {
        username: username,
        password: password,
        mode: mode
    };
    console.log(data)
    // 向后端发送数据
    fetch('/chat', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }).then(response => {
        // 这里可以处理响应
    }).catch(error => {
        console.error('Error:', error);
    });
});
