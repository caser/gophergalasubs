var Login = React.createClass({
  render: function() {
    return (<div id="login-portal">
            <img src="build/images/fancy-gopher.jpg" width="100px" />
            <h1>Login To Submit Your Gopher Gala 2016 Top 5</h1>
            <a href="/login" className="btn btn-green">Login With Github</a>
            </div>
           );
  }
});

module.exports = Login;

