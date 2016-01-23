var RepoItem = require('./RepoItem.jsx')
var Top5 = require('./Top5.jsx')
var Actions = require('../actions.js')

var Dashboard = React.createClass({

  vote: function(id) {
    Actions.submitVote(this.props.dispatch, id)
  },

  render: function() {
    var vote = this.vote;
    var repos;
    if(this.props.state.repos.length != 0) {
      repos = this.props.state.repos.map(function(repo){
        return <RepoItem key={repo.id} repo={repo} vote={function(){vote(repo.full_name)}} />
      })
      repos = (
        <ul id="submissions">
        {repos}
        </ul>
      )
    } else {
      repos = <img src="build/images/puff.svg" />
    }

    return (
      <div>
        <div id="passport" className="row">
          <div id="passport" className="col-md-8">
            <h1>Gopher Gala Voting</h1>
          </div>
          <div id="passport" className="col-md-4">
            {this.props.state.user.login}
             <img src={this.props.state.user.avatar_url} width="50px"/>
          </div>
        </div>
        <div className="row">
          <div className="col-md-8">
            <h2>Submissions</h2>
            {repos}
          </div>
          <div className="col-md-4">
            <h2>Your top 5</h2>
            <Top5 state={this.props.state} dispatch={this.props.dispatch}/>
          </div>
        </div>
      </div>
    );
  }
});

module.exports = Dashboard;
