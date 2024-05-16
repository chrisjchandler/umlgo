import os
from flask import Flask, request, render_template
import subprocess

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/analyze', methods=['POST'])
def analyze_go_code():
    # Receive Go code from the frontend
    go_code = request.form['go_code']

    # Path to the directory containing main.go and parser.go
    parser_dir = '.'

    # Execute the Go parser script as a subprocess, providing the code via stdin
    result = subprocess.run(['go', 'run', 'main.go'], input=go_code, cwd=parser_dir, capture_output=True, text=True)

    # Debugging: Print the result of the subprocess
    print("STDOUT:", result.stdout)
    print("STDERR:", result.stderr)

    # Check if the Go parser created the output file
    output_file_path = os.path.join(parser_dir, 'output.uml')
    if not os.path.exists(output_file_path):
        return "Error: output.uml file was not created by the Go parser.", 500

    # Read the generated UML code
    with open(output_file_path, 'r') as file:
        uml_code = file.read()

    # Remove temporary files
    os.remove(output_file_path)

    return render_template('result.html', uml_code=uml_code)

if __name__ == '__main__':
    app.run(debug=True)
