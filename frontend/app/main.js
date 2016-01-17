require('es6-promise').polyfill();
require('isomorphic-fetch');

var React = require('react');
var ReactDOM = require('react-dom');
var Redux = require('redux');

var ReduxThunk = require('redux-thunk');

var Util = require('./lib/util.js');
var Actions = require('./actions.js');

window.React = React; 

var reducer = function(state, action){
  console.log(action)
  switch (action.type) {
    case "repos_fetched":
      return Object.assign({}, state, {
        repos: action.repos
      })
    case "user_fetched":
      return Object.assign({}, state, {
        user: action.user
      })
    case "repo_voted":
      return Object.assign({}, state, {
        user: action.user
      })
    case "repo_vote_weighted":
      return Object.assign({}, state, {
        user: action.user
      })
    default:
      return state
  }
}

var enhancedStoreCreator = Redux.applyMiddleware(ReduxThunk)(Redux.createStore);

var defaultState = {
  loggedIn: !!Util.getCookie('token'),
  repos: [],
  pages: 0,
  user: {}
}

var store = enhancedStoreCreator(reducer, defaultState);

if(store.getState().loggedIn){
  store.dispatch(Actions.loadRepos)
  store.dispatch(Actions.loadUser)
}

var App = require('./components/App.jsx');

var render = function(){
  ReactDOM.render(<App state={store.getState()}/>, document.getElementById('app'));
}

store.subscribe(render);
render();
