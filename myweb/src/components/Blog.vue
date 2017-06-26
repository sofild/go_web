<template>
<div class="uk-grid" data-uk-grid-margin>
    <div class="uk-width-medium-3-4">
        <article class="uk-article" v-for="(item, index) in items">
            <h1 class="uk-article-title">
                <a href="layouts_post.html">{{item.title}}</a>
            </h1>
            <p class="uk-article-meta">{{item.addtime}} &nbsp;&nbsp;<a href="#">{{item.catename}}</a></p>
            <p>
                <a href="layouts_post.html"><img width="900" height="300" v-bind:src="item.pic" v-bind:alt="item.title"></a>
            </p>
            <p class="uk-article-meta">
                {{item.description}}
            </p>
        </article>

        <ul class="uk-pagination">
            <li class="uk-disabled" v-if="nowPage>1"><span><i class="uk-icon-angle-double-left" v-on:click="toPage(nowPage-1)"></i></span></li>
            <template v-for="p in totalPage">
              <template v-if="nowPage==p">
                <li class="uk-active"><span>{{nowPage}}</span></li>
              </template>
              <template v-else>
                <li><a href="javascript:void(0);" v-on:click="toPage(p)">{{p}}</a></li>
              </template>
            </template>
            <li v-if="nowPage<totalPage"><a href="javascript:void(0);" v-on:click="toPage(nowPage+1)"><i class="uk-icon-angle-double-right"></i></a></li>
        </ul>

    </div>

    <div class="uk-width-medium-1-4">
        <div class="uk-panel uk-panel-box uk-text-center" v-if="!!username">
            <img class="uk-border-circle" width="120" height="120" v-bind:src="(!!avatar) ? avatar : 'data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4NCjwhLS0gR2VuZXJhdG9yOiBBZG9iZSBJbGx1c3RyYXRvciAxNi4wLjQsIFNWRyBFeHBvcnQgUGx1Zy1JbiAuIFNWRyBWZXJzaW9uOiA2LjAwIEJ1aWxkIDApICAtLT4NCjwhRE9DVFlQRSBzdmcgUFVCTElDICItLy9XM0MvL0RURCBTVkcgMS4xLy9FTiIgImh0dHA6Ly93d3cudzMub3JnL0dyYXBoaWNzL1NWRy8xLjEvRFREL3N2ZzExLmR0ZCI+DQo8c3ZnIHZlcnNpb249IjEuMSIgaWQ9IkViZW5lXzEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHg9IjBweCIgeT0iMHB4Ig0KCSB3aWR0aD0iMTIwcHgiIGhlaWdodD0iMTIwcHgiIHZpZXdCb3g9IjAgMCAxMjAgMTIwIiBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCAxMjAgMTIwIiB4bWw6c3BhY2U9InByZXNlcnZlIj4NCjxyZWN0IGZpbGw9IiNGRkZGRkYiIHdpZHRoPSIxMjAiIGhlaWdodD0iMTIwIi8+DQo8Zz4NCgk8cGF0aCBmaWxsPSIjRTBFMEUwIiBkPSJNMTA5LjM1NCw5OS40NzhjLTAuNTAyLTIuODA2LTEuMTM4LTUuNDA0LTEuOTAzLTcuODAxYy0wLjc2Ny0yLjM5Ny0xLjc5Ny00LjczMi0zLjA5My03LjAxMQ0KCQljLTEuMjk0LTIuMjc2LTIuNzc4LTQuMjE3LTQuNDU1LTUuODIzYy0xLjY4MS0xLjYwNC0zLjcyOS0yLjg4Ny02LjE0OC0zLjg0NmMtMi40MjEtMC45NTgtNS4wOTQtMS40MzgtOC4wMTctMS40MzgNCgkJYy0wLjQzMSwwLTEuNDM3LDAuNTE2LTMuMDIsMS41NDVjLTEuNTgxLDEuMDMyLTMuMzY3LDIuMTgyLTUuMzU1LDMuNDVjLTEuOTksMS4yNzEtNC41NzgsMi40MjEtNy43NjUsMy40NTENCgkJQzY2LjQxLDgzLjAzNyw2My4yMSw4My41NTIsNjAsODMuNTUyYy0zLjIxMSwwLTYuNDEtMC41MTUtOS41OTgtMS41NDZjLTMuMTg4LTEuMDMtNS43NzctMi4xODEtNy43NjUtMy40NTENCgkJYy0xLjk5MS0xLjI2OS0zLjc3NC0yLjQxOC01LjM1NS0zLjQ1Yy0xLjU4Mi0xLjAyOS0yLjU4OC0xLjU0NS0zLjAyLTEuNTQ1Yy0yLjkyNiwwLTUuNTk4LDAuNDc5LTguMDE3LDEuNDM4DQoJCWMtMi40MiwwLjk1OS00LjQ3MSwyLjI0MS02LjE0NiwzLjg0NmMtMS42ODEsMS42MDYtMy4xNjQsMy41NDctNC40NTgsNS44MjNjLTEuMjk0LDIuMjc4LTIuMzI2LDQuNjEzLTMuMDkyLDcuMDExDQoJCWMtMC43NjcsMi4zOTYtMS40MDIsNC45OTUtMS45MDYsNy44MDFjLTAuNTAyLDIuODAzLTAuODM5LDUuNDE1LTEuMDA2LDcuODM1Yy0wLjE2OCwyLjQyMS0wLjI1Miw0LjkwMi0wLjI1Miw3LjQ0DQoJCWMwLDEuODg0LDAuMjA3LDMuNjI0LDAuNTgyLDUuMjQ3aDEwMC4wNjNjMC4zNzUtMS42MjMsMC41ODItMy4zNjMsMC41ODItNS4yNDdjMC0yLjUzOC0wLjA4NC01LjAyLTAuMjUzLTcuNDQNCgkJQzExMC4xOTIsMTA0Ljg5MywxMDkuODU3LDEwMi4yOCwxMDkuMzU0LDk5LjQ3OHoiLz4NCgk8cGF0aCBmaWxsPSIjRTBFMEUwIiBkPSJNNjAsNzguMTZjNy42MiwwLDE0LjEyNi0yLjY5NiwxOS41Mi04LjA4OGM1LjM5Mi01LjM5Myw4LjA4OC0xMS44OTgsOC4wODgtMTkuNTE5DQoJCXMtMi42OTYtMTQuMTI2LTguMDg4LTE5LjUxOUM3NC4xMjYsMjUuNjQzLDY3LjYyLDIyLjk0Niw2MCwyMi45NDZzLTE0LjEyOCwyLjY5Ny0xOS41MTksOC4wODkNCgkJYy01LjM5NCw1LjM5Mi04LjA4OSwxMS44OTctOC4wODksMTkuNTE5czIuNjk1LDE0LjEyNiw4LjA4OSwxOS41MTlDNDUuODcyLDc1LjQ2NCw1Mi4zOCw3OC4xNiw2MCw3OC4xNnoiLz4NCjwvZz4NCjwvc3ZnPg0K'" alt="">
            <h3>{{username}}</h3>
            <p>login time:{{logintime}}</p>
        </div>
        <div class="uk-panel">
            <h3 class="uk-panel-title">Time</h3>
            <ul class="uk-list uk-list-line">
                <li v-for="tItem in timeList"><a href="javascript:void(0);" v-bind:data="tItem.time">{{tItem.date}}</a></li>
            </ul>
        </div>
        <div class="uk-panel">
            <h3 class="uk-panel-title">Category</h3>
            <ul class="uk-list">
                <li v-for="cItem in cate"><a href="javascript:void(0);">{{ cItem.name }}</a></li>
            </ul>
        </div>
    </div>
</div>
</template>
<script>
  console.log(window.BlogList,window.BlogList.length);
  export default{
    name: 'Blog',
    data(){
      return {
        blogList: window.BlogList,
        items: window.BlogList.slice(0,2),
        cate: window.CateList,
        totalPage: Math.ceil(window.BlogList.length/2),
        nowPage: 1,
        username: window.ExtFunc.getCookie("username"),
        avatar: window.ExtFunc.getCookie("avatar"),
        logintime: window.ExtFunc.getCookie("logintime"),
      }
    },
    computed:{
        timeList:function(){
          var myDate = new Date();
          var year = myDate.getFullYear();
          var month = myDate.getMonth();
          month = month + 1;
          var data = [];
          for(var i=0; i<6; i++){
            var t = {};
            month = month - 1;
            if(month < 1){
              month = 12;
              year = year - 1;
            }
            var fMonth = month < 10 ? "0"+month+"" : ""+month+"";
            var date = ""+year+"/"+fMonth+"/01 00:00:00";
            var unixDate = new Date(Date.parse(date));
            var unixTime = unixDate.getTime();
            t.date = ""+year+"."+fMonth;
            t.time = unixTime;
            data.push(t);
          }
          return data;
        }
    },
    watch: {
      nowPage:function(p){
        return p;
      }
    },
    methods:{
      toPage: function(p){
        var start = (p-1)*2;
        var end = start + 2;
        this.nowPage = p;
        this.items = this.blogList.slice(start, end);
        console.log(this.items);
        console.log(this.nowPage);
      },
      selCate: function(c){
        var list = window.BlogList;
        var data = [];
        for(var i in list){
          if(list[i]["cateid"]==c){
            data.push(list[i]);
          }
        }
        this.blogList = data;
        this.items = data.slice(0, 2);
        this.totalPage = Math.ceil(this.blogList.length/2);
        this.nowPage = 1;
      },
      selTime: function(t){
        var list = window.BlogList;
        var data = [];
        for(var i in list){
          if(parseInt(list[i]["addtime"]) > t){
            data.push(list[i]);
          }
        }
        this.blogList = data;
        this.items = data.slice(0, 2);
        this.totalPage = Math.ceil(this.blogList.length/2);
        this.nowPage = 1;
      }
    }
  }
</script>
