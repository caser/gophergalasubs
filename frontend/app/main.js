var React = require('react');
var ReactDOM = require('react-dom');
var Redux = require('redux');
var Util = require('./lib/util.js');

window.React = React; 

var reducer = function(state, action){
  return state;
}

var store = Redux.createStore(reducer, {
  token: Util.getCookie('token')
});

var App = require('./components/App.jsx');

var render = function(){
  ReactDOM.render(<App state={store.getState()}/>, document.getElementById('app'));
}

store.subscribe('render');
render();
