<template>
  <div>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Users</span>
      </div>
      <el-table :data="users">
        <div v-for="name in userOptions" :key="name">
          <el-table-column :prop="name" :label="name"> </el-table-column>
        </div>
      </el-table>
    </el-card>
    <br>

    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Summoneds</span>
      </div>
      <label for="name">Search by Name</label>
      <el-input type="text" v-model="filterName" id="name" ></el-input>
      <br />
      <label>Search by Type</label>
      <br>
      <el-select v-model="filterType">
        <el-option v-for="item in options" :key="item" :label="item" :value="item"></el-option>
      </el-select>
      <el-table :data="summoneds">
        <div v-for="name in summonedOptions" :key="name">
          <el-table-column :prop="name" :label="name"> </el-table-column>
        </div>
        <el-table-column fixed="right" label="owner">
          <template slot-scope="scope">
            <router-link :to="'/profile/' + scope.row.user_id">info</router-link>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <br>

    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Requests</span>
      </div>
      <el-table :data="requests">
        <div v-for="name in requestOptions" :key="name">
          <el-table-column :prop="name" :label="name"> </el-table-column>
        </div>
        <el-table-column fixed="right" label="taker">
          <template slot-scope="scope">
            <router-link :to="'/profile/' + scope.row.user_id">info</router-link>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <br>

    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Transactions</span>
      </div>
      <el-table :data="transactions">
        <div v-for="name in transactionOptions" :key="name">
          <el-table-column :prop="name" :label="name"> </el-table-column>
        </div>
      </el-table>
    </el-card>
    <br>

    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Profits</span>
      </div>
      <label for="startTime">Start time</label>
      <el-input type="date" v-model="startTime" id="startTime" ></el-input>
      <br>
      <label for="endTime">End time</label>
      <el-input type="date" v-model="endTime" id="endTime" ></el-input>
      <br>
      <label>Summoned Type</label>
      <br>
      <el-select v-model="filterTypeProfit">
        <el-option v-for="item in options" :key="item" :label="item" :value="item"></el-option>
      </el-select>
      <br>
      <label>City</label>
      <br>
      <el-select v-model="filterCity">
        <el-option v-for="item in cityOptions" :key="item" :label="item" :value="item"></el-option>
      </el-select>
      <br>
      <div class="small">
        <chart :chart-data="profitChartData"></chart>
      </div>
      <el-table :data="filterProfits">
        <div v-for="name in profitOptions" :key="name">
          <el-table-column :prop="name" :label="name" sortable> </el-table-column>
        </div>
      </el-table>
    </el-card>
    <br>

    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>累计收益</span>
      </div>
    {{sum_profits}}
    </el-card>
  </div>
</template>

<script>
export default {
  components: {
    chart: () => import("../components/LineChart"),
  },
  data() {
    return {
      summonedOptions: [
        "status",
        "ID",
        "CreatedAt",
        "UpdatedAt",
        "user_id",
        "type",
        "name",
        "desc",
        "people",
        "ddl",
        "Img",
      ],
      userOptions: [
        "phone",
        "ID",
        "CreatedAt",
        "UpdatedAt",
        "username",
        "password",
      ],
      requestOptions: [
        "status",
        "ID",
        "CreatedAt",
        "UpdatedAt",
        "summoned_id",
        "user_id",
        "desc",
      ],
      transactionOptions: [
        "taker_cost",
        "ID",
        "CreatedAt",
        "UpdatedAt",
        "owner_id",
        "taker_id",
        "owner_cost",
      ],
      profitOptions: [
        "cost",
        "date",
        "city",
        "summoned_type",
        "count",
      ],
      cityOptions: [
        "",
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
      ],
      options: [
        "",
        "技术交流",
        "学业探讨",
        "社会实践",
        "公益志愿者",
        "游玩",
        "其它",
      ],
      allSummoneds: [],
      users: [],
      requests: [],
      transactions: [],
      profits: [],
      filterType: "",
      filterName: "",
      startTime: "",
      endTime: "",
      filterTypeProfit: "",
      filterCity: "",
    };
  },
  computed: {
    summoneds() {
      if (this.filterName == "" && this.filterType == "")
        return this.allSummoneds;
      if (this.filterName == "")
        return this.allSummoneds.filter((s) => s.type == this.filterType);
      if (this.filterType == "")
        return this.allSummoneds.filter((s) =>
          s.name.includes(this.filterName)
        );
      return this.allSummoneds
        .filter((s) => s.name.includes(this.filterName))
        .filter((s) => s.type == this.filterType);
    },
    filterProfits() {
      var res = this.profits
      if (this.startTime != "")
        res = res.filter((s) => new Date(s.date) >= new Date(this.startTime))
      if (this.endTime != "")
        res = res.filter((s) => new Date(s.date) <= new Date(this.endTime))
      if (this.filterTypeProfit != "")
        res = res.filter((s) => s.summoned_type == this.filterTypeProfit)
      if (this.filterCity != "")
        res = res.filter((s) => s.city == this.filterCity)
      return res      
    },
    sum_profits() {
      let res = 0
      this.transactions.forEach(t => res += t.owner_cost + t.taker_cost);
      return res
    },
    profitChartData() {
      if(this.filterProfits.length == 0)
        return {labels: [],datasets: []}
      var startTime, endTime
      if(this.startTime != ""){
        startTime = new Date(this.startTime)
      }else{
        startTime = new Date(this.filterProfits[0].date)
        for(let p=1; p < this.filterProfits.length; p++){
          if(new Date(this.filterProfits[p].date) < startTime)startTime = new Date(this.filterProfits[p].date)
        }
      }
      if(this.endTime != ""){
        endTime = new Date(this.endTime)
      }else{
        endTime = new Date(this.filterProfits[0].date)
        for(let p=1; p < this.filterProfits.length; p++){
          if(new Date(this.filterProfits[p].date) > endTime)endTime = new Date(this.filterProfits[p].date)
        }
      }
      endTime = new Date(endTime.setMonth(endTime.getMonth()+1))
      var labelsList = []
      var accCount = []
      var accCost = []
      for(var time=startTime; time<=endTime; time=new Date(time.setMonth(time.getMonth()+1))){
        labelsList.push(time.toLocaleDateString())
        let tmpCount = 0
        let tmpCost = 0
        for(let p=0; p < this.filterProfits.length; p++){
          if(new Date(this.filterProfits[p].date) <= time){
            tmpCount += this.filterProfits[p].count
            tmpCost += this.filterProfits[p].cost
          }
        }
        accCount.push(tmpCount)
        accCost.push(tmpCost)
      }

      var countData = []
      countData.push(accCount[0])
      var costData = []
      costData.push(accCost[0])
      for(let i=1;i<accCount.length;i++){
        countData[i] = accCount[i] - accCount[i-1]
      }
      for(let i=1;i<accCost.length;i++){
        costData[i] = accCost[i] - accCost[i-1]
      }
      
      return {
        labels: labelsList,
        datasets: [
          {
            label: 'Count',
            backgroundColor: '#f87979',
            data: countData
          },
          {
            label: 'Cost',
            backgroundColor: '#007979',
            data: costData
          }
        ]
      }
    }
  },
  mounted() {
    let that = this;
    this.$axios
      .get("summoned")
      .then(function (response) {
        that.allSummoneds = response.data;
      })
      .catch(function (response) {
        console.log(response);
      });
    this.$axios
      .get("users")
      .then(function (response) {
        that.users = response.data;
      })
      .catch(function (response) {
        console.log(response);
      });
    this.$axios
      .get("requestsAll")
      .then(function (response) {
        that.requests = response.data;
      })
      .catch(function (response) {
        console.log(response);
      });
    this.$axios
      .get("transaction")
      .then(function (response) {
        that.transactions = response.data;
      })
      .catch(function (response) {
        console.log(response);
      });
    this.$axios
      .get("profit")
      .then(function (response) {
        that.profits = response.data;
      })
      .catch(function (response) {
        console.log(response);
      });
  },
};
</script>

<style>
  .small {
    max-width: 400px;
  }
</style>