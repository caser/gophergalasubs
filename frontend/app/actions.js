var Util = require('./lib/util.js');

var Actions = {
  loadRepos: function(dispatch) {
    return fetch('/repos?token='+Util.getCookie('token')).then(function(resp) {
      return resp.json();
    }).then(function(json){
      dispatch({
        type:"repos_fetched",
        repos: json
      })
    })
  },

  loadUser: function(dispatch) {
    return fetch('/user?token='+Util.getCookie('token')).then(function(resp) {
      return resp.json();
    }).then(function(json){
      dispatch({
        type:"user_fetched",
        user: json
      })
    })
  }
}

module.exports = Actions
