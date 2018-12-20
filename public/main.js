var app = new Vue({
  el: '#app',
  data: {
    message: 'Hello Vue!',
    users: []
  },
  mounted() {
    fetch('/users')
        .then(resp => console.log('users', resp.data))
        .catch(err => console.error(err))
  }
});