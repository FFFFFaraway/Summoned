<template>
  <div>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Signup</span>
      </div>
    <label for="username">Username: </label>
    <el-input type="text" v-model="username" id="username" ></el-input>
    <label for="password">Password: </label>
    <el-input type="password" v-model="password" id="password" show-password></el-input>

    <label for="name">True Name: </label>
    <el-input type="text" v-model="name" id="name" ></el-input>
    <label>Type of certificate: </label>
    <br>
    <el-select v-model="numberType">
      <el-option v-for="item in option" :key="item" :label="item" :value="item"></el-option>
    </el-select>
    <br>
    <label for="number">Number of certificate: </label>
    <el-input type="number" v-model="number" id="number" ></el-input>
    <label for="phone">Phone: </label>
    <el-input type="number" v-model="phone" id="phone" ></el-input>
    <label for="introduction">Introduction(optional): </label>
    <el-input type="text" v-model="introduction" id="introduction" ></el-input>
    <label>Province: </label>
    <br>
    <el-select v-model="city">
      <el-option v-for="item in cityOption" :key="item" :label="item" :value="item"></el-option>
    </el-select>
    <br>
    <el-button type="primary" @click="submit">Sign up</el-button>
    <el-button type="primary" @click="reset">Reset</el-button>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      password: "",
      name: "",
      numberType: "",
      number: "",
      phone: "",
      introduction: "",
      city: "",
      option: [
        "身份证",
        "学生证"
      ],
      cityOption: [
        "北京市",
        "天津市",
        "上海市",
        "重庆市",
        "河北省",
        "山西省",
        "辽宁省", 
        "吉林省",
        "黑龙江省",
        "江苏省",
        "浙江省",
        "安徽省",
        "福建省",
        "江西省",
        "山东省",
        "河南省",
        "湖北省",
        "湖南省",
        "广东省",
        "海南省",
        "四川省",
        "贵州省",
        "云南省",
        "陕西省",
        "甘肃省",
        "青海省",
        "台湾省",
        "内蒙古自治区",
        "广西壮族自治区",
        "西藏自治区",
        "宁夏回族自治区",
        "新疆维吾尔自治区",
        "香港特别行政区",
        "澳门特别行政区",
      ]
    };
  },
  methods: {
    submit() {
      if (this.username.length == 0) {
        alert("输入用户名")
        return
      }
      if (this.password.length < 6) {
        alert("密码最短6位")
        return
      }
      var numberCnt = 0
      var haveUpLetter = false
      var haveLowLetter = false
      for(var i=0; i < this.password.length; i++) {
        if (this.password[i] >= '0' && this.password[i] <= '9')numberCnt += 1
        if (this.password[i] >= 'A' && this.password[i] <= 'Z')haveUpLetter = true
        if (this.password[i] >= 'a' && this.password[i] <= 'z')haveLowLetter = true
      }
      if (numberCnt < 2 || !haveUpLetter || !haveLowLetter){
        alert("密码至少两个数字，并且有大小写")
        return
      }
      if (this.name.length == 0) {
        alert("输入姓名")
        return
      }
      if (this.numberType.length == 0) {
        alert("请选择证件类型")
        return
      }
      if (this.number.length == 0) {
        alert("请输入证件号码")
        return
      }
      if (this.phone.length != 11) {
        alert("请输入11位电话号码")
        return
      }
      if (this.city.length == 0) {
        alert("请选择省")
        return
      }

      var formData = new FormData();
      formData.append("username", this.username);
      formData.append("password", this.password);
      formData.append("name", this.name);
      formData.append("number_type", this.numberType);
      formData.append("number", this.number);
      formData.append("phone", this.phone);
      formData.append("introduction", this.introduction);
      formData.append("city", this.city);
      let that = this;
      this.$axios
        .post("signup", formData, {
          withCredentials: true,
        })
        .then(function (response) {
          alert(response.data.message);
          if (response.data.message == "successfully sign up")
            that.$router.push("login");
        })
        .catch(function (response) {
          console.log(response);
        });
    },
    reset() {
      this.username = ""
      this.password = ""
      this.name = ""
      this.numberType = ""
      this.number = ""
      this.phone = ""
      this.introduction = ""
      this.city = ""
    },
  },
};
</script>