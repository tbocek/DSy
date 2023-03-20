/*

1) Navigate to the hydration directory in your terminal and run npm init -y to create a package.json file.
2) Install the required dependencies using the following command:
npm install vue@2 vue-server-renderer express axios
3) Ensure that the Go API server is running (use go run api.go if it's not).
4) Start the Node.js server by running node server.js in the hydration directory.
5) Now, you can open your browser and navigate to http://localhost:8081.
 */

// server.js
const Vue = require("vue");
const express = require("express");
const serverRenderer = require("vue-server-renderer").createRenderer();
const axios = require("axios");

const app = express();

app.get("/", async (req, res) => {
    try {
        const response = await axios.get("http://localhost:8082/api");
        const text = response.data.text;

        const vm = new Vue({
            template: `<div><h1>{{ text }}</h1></div>`,
            data: {
                text,
            },
        });

        serverRenderer.renderToString(vm, (err, html) => {
            if (err) {
                res.status(500).send("Internal Server Error");
                console.error(err);
                return;
            }

            const finalHtml = `
        <!DOCTYPE html>
        <html lang="en">
          <head>
            <meta charset="utf-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Vue SSR Hydration</title>
          </head>
          <body>
            <div id="app">${html}</div>
            <script>window.__INITIAL_STATE__ = ${JSON.stringify({ text })}</script>
            <script src="/client.js"></script>
          </body>
        </html>
      `;

            res.send(finalHtml);
        });
    } catch (error) {
        console.error("Error fetching data:", error);
        res.status(500).send("Internal Server Error");
    }
});

app.use(express.static("public"));

const port = 8081;
app.listen(port, () => {
    console.log(`Server is listening on port ${port}`);
});