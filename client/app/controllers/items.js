import Ember from 'ember';

export default Ember.Controller.extend({
  inputIncomplete: Ember.computed.not('inputComplete'),
  inputComplete: Ember.computed.and('code', 'payload', 'type'),
  actions:{
    submit(){
      var item = this.store.createRecord('item',
        this.getProperties('code', 'payload', 'type'));
      item.save().then(() => {
        this.setProperties({
          'code': null,
          'payload': null,
          'type': null
        });
      });
    }
  }
});
