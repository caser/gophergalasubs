var Actions = require('../actions.js');

var RepoItem = React.createClass({

  hasVoted: function(){
    var votes = [
      this.props.user.vote1,
      this.props.user.vote2,
      this.props.user.vote3,
      this.props.user.vote4,
      this.props.user.vote5
    ];
    
    return votes.indexOf(this.props.repo.id) >= 0;
  },

  canVote: function(){
    return  this.props.user.vote1 &&
      this.props.user.vote2 &&
      this.props.user.vote3 &&
      this.props.user.vote4 &&
      this.props.user.vote5
  },
  
  remove: function(){
    Actions.removeVote(this.props.dispatch, this.props.user, this.props.repo.id)
  },
  
  render: function(){
    var button;
    
    if (this.hasVoted()) {
      button = <button className="btn btn-default" onClick={this.remove}>
              Unvote <span className="glyphicon glyphicon-remove" aria-hidden="true"></span>
            </button>
    } else {
      button = <button disabled={this.canVote()} className="btn btn-green" onClick={this.props.vote}>
              Vote <span className="glyphicon glyphicon-ok" aria-hidden="true"></span>
            </button>
    }
    
    return (<li className="repo">
            <div className="details">
              <h3 className="name">
                {this.props.repo.name}
                <span className="stargazers" title={this.props.repo.stargazers_count + " Stargazers"}>
                  <span className="glyphicon glyphicon-star" aria-hidden="true"></span>
                  {this.props.repo.stargazers_count} 
                </span>
              </h3>
              <p className="description">{this.props.repo.description || "This submission has no description."}</p>
            </div>
            {button}
            </li>)
  }

});

module.exports = RepoItem;
