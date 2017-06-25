<template>
  <div>
    <div class="uk-grid" data-uk-grid-margin>
        <div class="uk-width-1-1">
            <h1 class="uk-heading-large">{{msg}}</h1>
            <p class="uk-text-large"><label v-for="item in cate">{{item.name}} </label> 各种图片应有尽有...</p>
        </div>
    </div>

    <div class="uk-grid" data-uk-grid-margin>
        <div class="uk-width-1-1">

            <ul class="uk-subnav uk-subnav-pill" data-uk-switcher="{connect:'#switcher-content'}">
                <li class="uk-active"><a href="javascript:void(0);" v-on:click="selCate(0)">All</a></li>
                <li v-for="cItem in cate"><a href="javascript:void(0);" v-on:click="selCate(cItem.id)">{{cItem.name}}</a></li>
            </ul>

            <ul id="switcher-content" class="uk-switcher">

                <li class="uk-active">
                    <div class="uk-grid" data-uk-grid-margin v-for="i in 3">
                        <div class="uk-width-medium-1-3" v-for="(item, index) in items.slice((i-1)*3, i*3)">
                            <div class="uk-thumbnail uk-overlay-hover" data-uk-modal="{target:'#modal-1'}" >
                                <figure class="uk-overlay">
                                    <img width="600" height="400" v-bind:src="item.image" alt="">
                                    <figcaption class="uk-overlay-panel uk-overlay-icon uk-overlay-background uk-overlay-fade"></figcaption>
                                    <a class="uk-position-cover" href="#"></a>
                                </figure>
                            </div>
                            <div id="modal-1" class="uk-modal">
                                <div class="uk-modal-dialog uk-modal-dialog-lightbox">
                                    <a href="" class="uk-modal-close uk-close uk-close-alt"></a>
                                    <img width="600" height="400" v-bind:src="item.image" alt="">
                                </div>
                            </div>
                        </div>
                     </div>
                </li>
            </ul>

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
    </div>
</div>
</template>
<script>
  export default{
    name:"Photo",
    data(){
      return {
        photoList: window.PhotoList,
        items: window.PhotoList.slice(0, 9),
        totalPage: Math.ceil(window.PhotoList.length/9),
        nowPage: 1,
        cate: window.CateList,
        msg: '图片中心'
      }
    },
    watch: {
       items: function(){
          return this.photoList.slice(0, 9);
       },
       totalPage: function(){
          return Math.ceil(this.photoList.length/9);
       }
    },
    methods: {
      page: p => {
        var start = (p-1)*9;
        var end = start + 9;
        this.nowPage = p;
        this.items = this.photoList.slice(start, end);
      },
      selCate: function(c){
        if(c > 0){
          var data = [];
          for(var i in window.PhotoList){
              if(window.PhotoList[i]['cateid']==c){
                data.push(window.PhotoList[i]);
              }
          }
          this.photoList = data;
          console.log(data);
          this.items = this.photoList.slice(0, 9);
          console.log(this.items);
          this.totalPage = Math.ceil(this.photoList.length/9);
          this.nowPage = 1;
          console.log(this.totalPage, this.nowPage);
        }
        else{
          this.photoList = window.PhotoList;
          this.items = this.photoList.slice(0, 9);
          this.totalPage = Math.ceil(this.photoList.length/9);
          this.nowPage = 1;
        }
      }
    }
  }
</script>
