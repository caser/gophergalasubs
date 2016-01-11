var React = require('react');
var ReactDOM = require('react-dom');
var Redux = require('redux');

window.React = React; 

var store = Redux.createStore();
var App = require('./components/App.jsx');

var render = function(){
  ReactDOM.render(<App/>, document.getElementById('app'));
}

// store.subscribe('render');
render();
