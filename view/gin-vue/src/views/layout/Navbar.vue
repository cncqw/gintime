<template>
    <div>
  <b-navbar toggleable="lg" type="dark" variant="info">

    <b-container>
    <b-navbar-brand href="#">Alldu</b-navbar-brand>

    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

    <b-collapse id="nav-collapse" is-nav>
      <b-navbar-nav>
        <b-nav-item href="#">首页</b-nav-item>
        <b-nav-item href="#" disabled>发现</b-nav-item>
      </b-navbar-nav>

      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">
        <b-nav-item-dropdown right  v-if="userInfo">
          <template v-slot:button-content>
              <em>{{userInfo.name}}</em>
          </template>
               <b-dropdown-item @click="$router.replace({name:'login'})">个人主页</b-dropdown-item>
               <b-dropdown-item @click="$router.replace({name:'login'})">注销</b-dropdown-item>
        </b-nav-item-dropdown>
        <template v-if="!userInfo">
            <b-nav-item v-if="$route.name != 'login'" @click="$router.replace({name:'login'})">登录</b-nav-item>
            <b-nav-item v-if="$route.name != 'register'"  @click="$router.replace({name:'register'})">注册</b-nav-item>
        </template>
      </b-navbar-nav>
    </b-collapse>
    </b-container>
  </b-navbar>
</div>
</template>

<script>
import { mapState } from 'vuex';

export default {
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),

  // computed: {
  //   userInfo() {
  //     // return JSON.parse(storageService.get(storageService.USER_INFO));
  //     return this.$store.state.userModule.userInfo;
  //   },
  // },
};
</script>
