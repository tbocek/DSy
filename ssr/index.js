const express = require('express');
const app = express();
const port = 3000;

// Route to handle dynamic HTML generation
app.get('/', (req, res) => {
    // Data that will be used to generate dynamic content
    const data = {
        title: 'Server-Side Rendering Example',
        message: 'Hello from the server!',
        currentTime: new Date().toLocaleTimeString()
    };

    // Generate dynamic HTML using template literals
    const html = `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>${data.title}</title>
        </head>
        <body>
            <h1>${data.message}</h1>
            <p>The current time is: ${data.currentTime}</p>
            <a href="/about">Go to About Page</a>
        </body>
        </html>
    `;

    // Send the generated HTML as the response
    res.send(html);
});

// Route to handle the about page
app.get('/about', (req, res) => {
    const data = {
        title: 'About Us',
        message: 'This is the about page.',
        currentTime: new Date().toLocaleTimeString()
    };

    // Generate dynamic HTML for the about page
    const html = `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>${data.title}</title>
        </head>
        <body>
            <h1>${data.message}</h1>
            <p>The current time is: ${data.currentTime}</p>
            <a href="/">Go to Home Page</a>
        </body>
        </html>
    `;

    // Send the generated HTML as the response
    res.send(html);
});

app.get('/api', (req, res) => {
    // Generate dynamic HTML using template literals
    const html = `
        {"hallo":"test"}
    `;

    // Send the generated HTML as the response
    res.send(html);
});

// Start the server
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});
