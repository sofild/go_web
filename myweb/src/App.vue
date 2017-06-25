<template>
  <div id="app" class="uk-container uk-container-center uk-margin-top uk-margin-large-bottom">
    <nav class="uk-navbar uk-margin-large-bottom">
      <a class="uk-navbar-brand uk-hidden-small" href="layouts_frontpage.html">Brand</a>
      <ul class="uk-navbar-nav uk-hidden-small">
        <li class="uk-active">
          <router-link to="/">首页</router-link>
        </li>
        <li>
          <router-link to="/photo">图片</router-link>
        </li>
        <li>
          <router-link to="/blog">文章</router-link>
        </li>
        <li>
          <a href="/login" target="_blank">登录</a>
        </li>
      </ul>
      <a href="#offcanvas" class="uk-navbar-toggle uk-visible-small" data-uk-offcanvas></a>
      <div class="uk-navbar-brand uk-navbar-center uk-visible-small">Brand</div>
    </nav>

    <router-view></router-view>
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
