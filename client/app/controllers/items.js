import Ember from 'ember';

export default Ember.Controller.extend({
  actions:{
    submit(){
      this.store.createRecord('item', 
        this.getProperties('code', 'payload', 'type'));
    }
  }
});
