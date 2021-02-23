<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="登录">
          <b-form>
            <b-form-group label="手机号">
              <b-form-input v-model="$v.user.telephone.$model"
              type="number"
              placeholder="输入手机号"
              required :state="validateState('telephone')">
              </b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')">
                  手机号格式不正确
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input v-model="$v.user.password.$model"
              type="password"
              placeholder="请输入密码"
             :state="validateState('password')">
              </b-form-input>
               <b-form-invalid-feedback :state="validateState('password')">
                 密码必须不少于6位
              </b-form-invalid-feedback>
            </b-form-group>

            <b-button block variant="primary" @click="register">登录</b-button>
            <b-button @click="$router.replace({name:'register'})" block variant="outline-secondary">注册</b-button>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>

import { required, minLength } from 'vuelidate/lib/validators';
import customValidator from '@/helper/validator';

export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
      },
      validation: null,
    };
  },
  validations: {
    user: {
      telephone: {
        required,
        telephone: customValidator.telephoneValidator,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    validateState(name) {
      // 这里是es6 解构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      if (this.user.telephone.length !== 11) {
        this.validate = false;
        return;
      }
      this.validate = true;
      console.log('r');
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
