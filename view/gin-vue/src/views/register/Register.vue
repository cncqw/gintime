<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="注册">
          <b-form>
            <b-form-group label="手机号">
              <b-form-input v-model="$v.user.telephone.$model"
              type="number"
              placeholder="输入手机号"
              required>
              </b-form-input>
              <b-form-invalid-feedback :state="validation">
                  手机号必须为11位
              </b-form-invalid-feedback>
            </b-form-group>
               <b-form-group label="昵称">
              <b-form-input v-model="$v.user.name.$model" type="text" placeholder="输入您的名称（选填）">
              </b-form-input>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input v-model="$v.user.password.$model"
              type="password"
              placeholder="请输入密码"
              required>
              </b-form-input>
            </b-form-group>

            <b-button block variant="primary" @click="register">确认</b-button>
            <b-button type="reset" block variant="outline-secondary">重置</b-button>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>

import { required, minLength, maxLength } from 'vuelidate/lib/validators';

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
        minLength: minLength(11),
        maxLength: maxLength(11),
      },
      name: {},
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
