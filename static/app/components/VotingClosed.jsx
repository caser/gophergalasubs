var VotingClosed = React.createClass({

  render: function() {
    var vote = this.vote;
    var user = this.props.state.user;
    var dispatch = this.props.dispatch;
    var repos;
    
    return (
      <div className="app">
        <div className="row">
          <div className="col-md-8">
            <img className="logo" src="build/images/fancy-gopher.jpg" width="100px" />
            <h1>Gopher Gala Votetastic</h1>
          </div>
          <div id="passport" className="col-md-4">
           {this.props.state.user.login} | 
             <span><a href="/logout">Logout</a></span>
            <img src={this.props.state.user.avatar_url} width="50px" />
          </div>
        </div>
        <div className="row">
          <div className="col-md-12">
            <h2>Submissions</h2>
            <p>Voting is now closed and we're tallying up the results. Good luck!</p>
          </div>
        </div>
      </div>
    );
  }
});

module.exports = VotingClosed
