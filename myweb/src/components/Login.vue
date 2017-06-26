<template>
<div class="uk-height-1-1">
  <div class="uk-vertical-align uk-text-center uk-height-1-1">
    <div class="uk-vertical-align-middle" style="width: 250px;">

      <!--<img class="uk-margin-bottom" width="140" height="120" src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4NCjwhLS0gR2VuZXJhdG9yOiBBZG9iZSBJbGx1c3RyYXRvciAxNi4wLjQsIFNWRyBFeHBvcnQgUGx1Zy1JbiAuIFNWRyBWZXJzaW9uOiA2LjAwIEJ1aWxkIDApICAtLT4NCjwhRE9DVFlQRSBzdmcgUFVCTElDICItLy9XM0MvL0RURCBTVkcgMS4xLy9FTiIgImh0dHA6Ly93d3cudzMub3JnL0dyYXBoaWNzL1NWRy8xLjEvRFREL3N2ZzExLmR0ZCI+DQo8c3ZnIHZlcnNpb249IjEuMSIgaWQ9IkViZW5lXzEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHg9IjBweCIgeT0iMHB4Ig0KCSB3aWR0aD0iMTQwcHgiIGhlaWdodD0iMTIwcHgiIHZpZXdCb3g9Ii0yOS41IDI3NS41IDE0MCAxMjAiIGVuYWJsZS1iYWNrZ3JvdW5kPSJuZXcgLTI5LjUgMjc1LjUgMTQwIDEyMCIgeG1sOnNwYWNlPSJwcmVzZXJ2ZSI+DQo8ZyBvcGFjaXR5PSIwLjciPg0KCTxwYXRoIGZpbGw9IiNEOEQ4RDgiIGQ9Ik0tNi4zMzMsMjk4LjY1NHY3My42OTFoOTMuNjY3di03My42OTFILTYuMzMzeiBNNzkuNzg4LDM2NC4zNTVIMS42NTZ2LTU3LjcwOWg3OC4xMzJWMzY0LjM1NXoiLz4NCgk8cG9seWdvbiBmaWxsPSIjRDhEOEQ4IiBwb2ludHM9IjUuODYsMzU4LjE0MSAyMS45NjIsMzQxLjIxNiAyNy45OTUsMzQzLjgyNyA0Ny4wMzIsMzIzLjU2MSA1NC41MjQsMzMyLjUyMyA1Ny45MDUsMzMwLjQ4IA0KCQk3Ni4yMDMsMzU4LjE0MSAJIi8+DQoJPGNpcmNsZSBmaWxsPSIjRDhEOEQ4IiBjeD0iMjQuNDYyIiBjeT0iMzIxLjMyMSIgcj0iNy4wMzQiLz4NCjwvZz4NCjwvc3ZnPg0K" alt="">-->

      <form class="uk-panel uk-panel-box uk-form" method="post" action="/doLogin" id="loginForm">
        <div class="uk-form-row">
          <input class="uk-width-1-1 uk-form-large" type="text" v-model="items.username" name="username" placeholder="Username">
        </div>
        <div class="uk-form-row">
          <input class="uk-width-1-1 uk-form-large" type="text" v-model="items.password" name="password" placeholder="Password">
        </div>
        <div class="uk-form-row">
          <a class="uk-width-1-1 uk-button uk-button-primary uk-button-large" href="javascript:void(0);" v-on:click="submitForm">Login</a>
        </div>
        <div class="uk-form-row uk-text-small">
          <label class="uk-float-left"><input type="checkbox"> Remember Me</label>
          <a class="uk-float-right uk-link uk-link-muted" href="#">Forgot Password?</a>
        </div>
      </form>

    </div>
  </div>
</div>
</template>
<script>
  export default{
    name: "Login",
    data(){
      return {
        items: {username:"", password:""}
      }
    },
    methods: {
      submitForm: function () {
        var username = this.items.username;
        var password = this.items.password;
        if (username == "" || password == "") {
          alert("请填写用户名和密码！");
          return false;
        }
        $.ajax({
          url: "/doLogin",
          type: "post",
          data: "username=" + username + "&password=" + password + "",
          dataType: "json",
          success: function (data) {
            if (data.err == 0) {
              window.ExtFunc.setCookie("username", data.username);
              window.ExtFunc.setCookie("avatar", data.avatar);
              window.ExtFunc.setCookie("logintime", data.logintime);
              window.location.href = "/";
            }
            else {
              alert(data.msg);
              return false;
            }
          }
        });
      }
    }
  }
</script>
