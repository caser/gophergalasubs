var Login = require('./Login.jsx');
var Dashboard = require('./Dashboard.jsx');

var App = React.createClass({
  render: function() {
    var body;
    if(this.props.state.token) {
      body = <Dashboard />;
    } else {
      body = <Login />;
    }
    return(
      <div>
      <h1>Gopher Gala Voting</h1>
      {body}
      </div>
    );
    console.log(this.props.state)
  }
});

module.exports = App;
