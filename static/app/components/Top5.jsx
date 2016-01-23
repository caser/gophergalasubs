var Actions = require('../actions.js');

var Top5 = React.createClass({

  render: function(){
    var state = this.props.state;

    if(this.props.state.repos.length == 0) {
      return null
    } else if (this.props.state.user.vote1 == null) {
      return (<p>You haven't voted yet</p>);
    }

    var user = state.user;
    var votes = [user.vote1, user.vote2, user.vote3, user.vote4, user.vote5];

    votes = votes.filter(function(n){
      return n != undefined;
    })

    votes = votes.map(function(n){
      var repo = state.repos.find(function(x){
        return x.id == n;
      })
      return (
      <li>
      {repo.name}
        <span className="glyphicon glyphicon-remove" aria-hidden="true"></span>
      </li>
      )
    });
    

    return(<ol id="top5">
           {votes}
    </ol>)
  }
});

module.exports = Top5;

