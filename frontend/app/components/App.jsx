var Actions = require('../actions.js');

var ReactRedux = require('react-redux');
var Login = require('./Login.jsx');
var Dashboard = require('./Dashboard.jsx');

var select = function(state){
  return state;
}

var App = React.createClass({
  render: function() {
    var body;
    if(this.props.state.loggedIn) {
      body = <Dashboard state={this.props.state} dispatch={this.props.dispatch}/>;
    } else {
      body = <Login />;
    }
    return(
      <div>
      {body}
      </div>
    );
  }
});

module.exports = ReactRedux.connect(select)(App);
