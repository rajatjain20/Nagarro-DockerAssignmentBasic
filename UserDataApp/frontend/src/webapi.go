// This file contains all the handle functions
// for this frontend application.

package main

import (
	"fmt"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	sEnv := envData.envName
	fmt.Fprintln(w, "Welcome to docker basic Assignment's frontend application running on "+sEnv)
}

// to add user, it will return a webpage of Add User.
func addUser(w http.ResponseWriter, r *http.Request) {
	server := "localhost:" + envData.backend_port

	w.Header().Set("Access-Control-Allow-Origin", "*") // Add CORS header
	fmt.Fprintf(w, `
    <html>
        <head>
            <title>User Data App - Add User</title>
            <style>
                table {
                    border-collapse: collapse;
                    width: auto;
                    margin: 20px 0;
                }
                table, th, td {
                    border: 1px solid black;
                }
                th, td {
                    padding: 8px;
                    text-align: left;
                }

                form { 
                    border: 2px solid black; 
                    padding: 10px; 
                    width: 300px; 
                    margin: 10px; 
                }

                .input-margin-ID { 
                    margin-left: 23px; 
                    margin-right: 0px; 
                }

                .input-margin-Name { 
                    margin-left: 20px; 
                    margin-right: 20px; 
                }

                /* Hide the spin buttons in WebKit browsers */
                input::-webkit-outer-spin-button,
                input::-webkit-inner-spin-button {
                    -webkit-appearance: none;
                    margin: 0;
                }

                /* Hide spin buttons in Firefox */
                input[type="number"] {
                    -moz-appearance: textfield;
                }
            </style>
            <script>
                function formatAddUserResponse(response) {
                    var table = '<table><tr><th>Rows Inserted</th><th>Message</th></tr>';
                    table += '<tr><td>' + response.RowsInserted + '</td><td>' + response.Message + '</td></tr>';
                    table += '</table>';
                    return table;
                }

                function submitForm(event) {
                    event.preventDefault();
                    var id = parseInt(document.getElementById('id').value);
                    var name = document.getElementById('name').value;
                    
                    var xhr = new XMLHttpRequest();
                    xhr.open('POST', 'http://`+server+`/addUser', true);
                    xhr.setRequestHeader('Content-Type', 'application/json');
                    xhr.onreadystatechange = function () {
                        if (xhr.readyState === 4) {
                            var responseElem = document.getElementById('response');
                            if (xhr.status === 200) {
                                var response = JSON.parse(xhr.responseText);
                                responseElem.innerHTML = formatAddUserResponse(response);
                            } else {
                                responseElem.innerHTML = 'Error: Status code - ' + xhr.status + ' - Unable to connect to backend server. ' + xhr.statusText;
                            }
                        }
                    };
                    var data = JSON.stringify({ "id": id, "name": name });
                    xhr.send(data);
                }

                function checkDBConnection() {
                    var xhr = new XMLHttpRequest();
                    xhr.open('GET', 'http://`+server+`/checkDB', true);
                    xhr.onreadystatechange = function () {
                        if (xhr.readyState === 4) {
                            if (xhr.status === 200) {
                                alert(xhr.responseText);
                            } else {
                                var err = 'Error: ' + xhr.status + ' - ' + xhr.statusText;
                                alert(err);
                            }
                        }
                    };
                    xhr.send();
                }
            </script>
        </head>
        <body>
            <!-- <button onclick="checkDBConnection()">Check Database Connection</button><br> -->
            <h2>User Data App - `+envData.envName+`</h2>
            <center>
            <h2>Add User Details</h2>
            <form onsubmit="submitForm(event)">
                <h3>Add User</h3>
                <label for="id">User ID:</label>
                <input type="number" id="id" name="id" class="input-margin-ID" required><br><br>
                <label for="name">User Name:</label>
                <input type="text" id="name" name="name" class="input-margin-Name" required><br><br>
                <button type="submit">Submit</button>
            </form>
            <p id="response"></p>
            </center>
        </body>
    </html>
    `)
}

// to get user data, it will return a webpage of Get User Info.
func getUserInfo(w http.ResponseWriter, r *http.Request) {
	server := "localhost:" + envData.backend_port

	w.Header().Set("Access-Control-Allow-Origin", "*") // Add CORS header

	fmt.Fprintf(w, `
    <html>
    <head>
        <title>User Data App - Get User</title>
        <style>
            table {
                border-collapse: collapse;
                width: auto;
                margin: 20px 0;
            }
            table, th, td {
                border: 1px solid black;
            }
            th, td {
                padding: 8px;
                text-align: left;
            }

            form { 
                border: 2px solid black; 
                padding: 20px; 
                width: 300px; 
                margin: 20px; 
            }

            .input-margin { 
                margin-left: 20px;  
                margin-right: 20px; 
            }
            
            /* Hide the spin buttons in WebKit browsers */
            input::-webkit-outer-spin-button,
            input::-webkit-inner-spin-button {
                -webkit-appearance: none;
                margin: 0;
            }

            /* Hide spin buttons in Firefox */
            input[type="number"] {
                -moz-appearance: textfield;
            }
        </style>
        <script>
            function formatTable(data) {
                if (!Array.isArray(data)) {
                    data = [data];
                }
                var table = '<table><tr><th>ID</th><th>Name</th><th>Message</th></tr>';
                for (var i = 0; i < data.length; i++) {
                    table += '<tr><td>' + data[i].ID + '</td><td>' + data[i].NAME + '</td><td>' + data[i].Message + '</td></tr>';
                }
                table += '</table>';
                return table;
            }

            function getUser() {
                var id = document.getElementById('id').value;
                var xhr = new XMLHttpRequest();
                xhr.open('GET', 'http://`+server+`/getUserInfo?id=' + id, true);
                xhr.onreadystatechange = function () {
                    if (xhr.readyState === 4) {
                        var responseElem = document.getElementById('response');
                        if (xhr.status === 200) {
                            var user = JSON.parse(xhr.responseText);
                            responseElem.innerHTML = formatTable(user);
                        } else {
                            'Error: Status code - ' + xhr.status + ' - Unable to connect to backend server. ' + xhr.statusText;
                        }
                    }
                };
                xhr.send();
            }

            function getAllUsers() {
                var xhr = new XMLHttpRequest();
                xhr.open('GET', 'http://`+server+`/getUserInfo', true);
                xhr.onreadystatechange = function () {
                    if (xhr.readyState === 4) {
                        var responseElem = document.getElementById('response');
                        if (xhr.status === 200) {
                            var users = JSON.parse(xhr.responseText);
                            responseElem.innerHTML = formatTable(users);
                        } else {
                            responseElem.innerHTML = 'Error: Status code - ' + xhr.status + ' - Unable to connect to backend server. ' + xhr.statusText;
                        }
                    }
                };
                xhr.send();
            }
        </script>
    </head>
    <body>
        <h2>User Data App - `+envData.envName+`</h2>
        <center>
        <h2>Get User Data</h2>
        <form onsubmit="event.preventDefault(); getUser();">
            <h3>Get User by ID</h3>
            <label for="id">User ID:</label>
            <input type="number" id="id" name="id" class="input-margin" required><br><br>
            <button type="submit">Get User</button>
        </form>
        <br>
        <button onclick="getAllUsers()">Get All Users</button>
        <p id="response"></p>
        </center>
    </body>
    </html>
    `)
}
