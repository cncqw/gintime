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
              required :state="validateState('telephone')">
              </b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')">
                  手机号格式不正确
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
             :state="validateState('password')">
              </b-form-input>
               <b-form-invalid-feedback :state="validateState('password')">
                 密码必须不少于6位
              </b-form-invalid-feedback>
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

import { required, minLength } from 'vuelidate/lib/validators';
import customValidator from '@/helper/validator';
import { mapActions } from 'vuex';

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
      name: {},
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    ...mapActions('userModule', { userRegister: 'register' }),

    validateState(name) {
      // 这里是es6 解构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      this.userRegister(this.user).then(() => {
        // 跳转主页
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '提示',
          variant: 'danger',
          solid: true,
          toaster: 'b-toaster-top-right',
        });
      });
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
