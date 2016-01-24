require('es6-promise').polyfill();
require('isomorphic-fetch');

var React = require('react');
var ReactDOM = require('react-dom');
var ReactRedux = require('react-redux');
var Redux = require('redux');
var _ = require('underscore');

var ReduxThunk = require('redux-thunk');

var Util = require('./lib/util.js');
var Actions = require('./actions.js');

window.React = React; 

var reducer = function(state, action){
  console.log(['ACTION', action.type, action])
  switch (action.type) {
    case "repos_fetched":
      return Object.assign({}, state, {
        repos: _.shuffle(action.repos)
      })
    case "user_fetched":
      return Object.assign({}, state, {
        user: action.user
      })
    case "user_cached":
      console.log(action)
      return Object.assign({}, state, {
        user: Object.assign({}, state.user, action.data)
      })
    default:
      return state
  }
}

var enhancedStoreCreator = Redux.applyMiddleware(ReduxThunk)(Redux.createStore);

var defaultState = {
  loggedIn: !!Util.getCookie('token'),
  repos: [],
  user: {}
}

var store = enhancedStoreCreator(reducer, defaultState);

if(store.getState().loggedIn){
  store.dispatch(Actions.loadRepos)
  store.dispatch(Actions.loadUser)
}

var App = require('./components/App.jsx');

var render = function(){
  ReactDOM.render(
    <ReactRedux.Provider store={store}>
      <App state={store.getState()}/>
    </ReactRedux.Provider>,
    document.getElementById('app')
  );
}

store.subscribe(render);
render();
