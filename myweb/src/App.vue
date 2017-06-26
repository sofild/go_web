<template>
  <div id="app" class="uk-container uk-container-center uk-margin-top uk-margin-large-bottom">
    <nav class="uk-navbar uk-margin-large-bottom">
      <a class="uk-navbar-brand uk-hidden-small" href="layouts_frontpage.html">Brand</a>
      <ul class="uk-navbar-nav uk-hidden-small">
        <li class="uk-active">
          <router-link to="/">HOME</router-link>
        </li>
        <li>
          <router-link to="/photo">GALLERY</router-link>
        </li>
        <li>
          <router-link to="/blog">ARTICLES</router-link>
        </li>
        <li>
          <router-link to="/login">LOGIN</router-link>
        </li>
      </ul>
      <a href="#offcanvas" class="uk-navbar-toggle uk-visible-small" data-uk-offcanvas></a>
      <div class="uk-navbar-brand uk-navbar-center uk-visible-small">Brand</div>
    </nav>
    <router-view></router-view>
  </div>
  <div id="offcanvas" class="uk-offcanvas">
    <div class="uk-offcanvas-bar">
      <ul class="uk-nav uk-nav-offcanvas">
        <li class="uk-active">
          <a href="/">HOME</a>
        </li>
        <li>
          <a href="/photo">GALLERY</a>
        </li>
        <li>
          <a href="/blog">ARTICLES</a>
        </li>
        <li>
          <a href="/login">LOGIN</a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'app',
  data(){
    return {
    }
  }
}
</script>
