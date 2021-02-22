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
           <b-nav-item @click="$router.replace({name:'login'})">登录</b-nav-item>
        <b-nav-item  @click="$router.replace({name:'register'})">注册</b-nav-item>
      </b-navbar-nav>
    </b-collapse>
    </b-container>
  </b-navbar>
</div>
</template>
