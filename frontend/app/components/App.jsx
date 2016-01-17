var Login = require('./Login.jsx');
var Dashboard = require('./Dashboard.jsx');

var App = React.createClass({
  render: function() {
    var body;
    if(this.props.state.loggedIn) {
      body = <Dashboard state={this.props.state}/>;
    } else {
      body = <Login />;
    }
    return(
      <div>
      <h1>Gopher Gala Voting</h1>
      {body}
      </div>
    );
  }
});

module.exports = App;
