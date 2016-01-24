var ReactDOM = require('react-dom');
var Actions = require('../actions.js');
var _ = require("underscore");

var Top5 = React.createClass({

  remove: function(id){
    Actions.removeVote(this.props.dispatch, this.props.state.user, id)
  },

  handleSortableUpdate: function(){
    var $node = $(ReactDOM.findDOMNode(this));
    var ids = $node.sortable('toArray', { attribute: 'data-id' });
    // We'll cancel the sortable change and let React reorder the DOM instead:
    $node.sortable('cancel');
    Actions.updateUser(this.props.dispatch, {
      'vote1': ids[0] ? parseInt(ids[0], 10) : null,
      'vote2': ids[1] ? parseInt(ids[1], 10) : null,
      'vote3': ids[2] ? parseInt(ids[2], 10) : null,
      'vote4': ids[3] ? parseInt(ids[3], 10) : null,
      'vote5': ids[4] ? parseInt(ids[4], 10) : null
    }, true)
  },

  componentDidUpdate() {
    $(ReactDOM.findDOMNode(this)).sortable({
      items: 'li',
      update: this.handleSortableUpdate
    });
  },

  render: function(){
    var state = this.props.state;
    var remove = this.remove;

    if(this.props.state.repos.length == 0) {
      return null
    } else if (this.props.state.user.vote1 == null) {
      return (<p>You haven't voted yet</p>);
    }

    var user = state.user;
    var votes = [user.vote1, user.vote2, user.vote3, user.vote4, user.vote5];

    votes = votes.filter(function(n){
      return n != undefined && n != 0;
    })

    votes = votes.map(function(n){
      var repo = state.repos.find(function(x){
        return x.id == n;
      })
      return (
        <li key={repo.id} data-id={repo.id} className="top-repo">
        <h3 className="name">{repo.name}</h3>
        <a className="btn btn-danger" onClick={
          function(){
            remove(repo.id)
          }
        }>
        <span className="glyphicon glyphicon-remove" aria-hidden="true"></span>
        </a>
      </li>
      )
    });
    

    return(<ul id="top5" className="sortable">
           {votes}
    </ul>)
  }
});

module.exports = Top5;

