// client.js
const Vue = require("vue");

const initialState = window.__INITIAL_STATE__ || {};

const vm = new Vue({
    el: "#app",
    template: `<div><h1>{{ text }}</h1></div>`,
    data: initialState,
});

delete window.__INITIAL_STATE__;