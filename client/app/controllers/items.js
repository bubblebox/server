import Ember from 'ember';

export default Ember.Controller.extend({
  actions:{
    submit(){
      var item = this.store.createRecord('item', 
        this.getProperties('code', 'payload', 'type'));
      item.save()
    }
  }
});
