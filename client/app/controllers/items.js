import Ember from 'ember';

export default Ember.Controller.extend({
  inputIncomplete: Ember.computed(function() {
    console.log('Make it work!');
    return true;
    
  }), 
  actions:{
    submit(){
      var item = this.store.createRecord('item', 
        this.getProperties('code', 'payload', 'type'));
      item.save()
    }
  }
});
