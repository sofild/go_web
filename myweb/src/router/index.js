import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Blog from '@/components/Blog'
import Photo from '@/components/Photo'
import Login from '@/components/Login'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    },
    {
      path: '/blog',
      name: 'Blog',
      component: Blog
    },
    {
      path: '/photo',
      name: 'Photo',
      component: Photo
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    }
  ]
})
