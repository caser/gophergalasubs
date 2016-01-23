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
  },

  submitVote: function(dispatch, id) {
    return fetch('/vote/' + id + '?token='+Util.getCookie('token'), {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        id: id
      })
    }).then(function(resp) {
      return resp.json();
    }).then(function(json){
      dispatch({
        type:"user_fetched",
        user: json
      })
    })
  },

  removeVote(dispatch, user, id) {
    votes = [user.vote1, user.vote2, user.vote3, user.vote4, user.vote5]
    votes = votes.filter(function(x){
      return x != id;
    });
    data = {
      vote1: votes[0],
      vote2: votes[1],
      vote3: votes[2],
      vote4: votes[3],
      vote5: votes[4]
    }
    return this.updateUser(dispatch, data);
  },

  updateUser: function(dispatch, data) {
    return fetch('/user?token='+Util.getCookie('token'), {
      method: 'PATCH',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    }).then(function(resp) {
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
