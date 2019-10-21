<template>
  <center>
    <h1>Member</h1>
    <form v-on:submit.prevent="submit">
        <input v-model="firstname" placeholder="Firstname" />
        <br/><br/>
        <input v-model="lastname" placeholder="Lastname" />
        <br/><br/>
        <input v-model="age" placeholder="Age" />
        <br/><br/>
        <select v-model="gender" >
            <option value disabled hidden>Select Gender</option>
            <option
            v-for="(gender,index) in genders"
            :key="index"
            v-bind:value="gender.gender_id"
            >{{gender.genderName}}</option>
        </select>
        <br/><br/>
        <button type="submit" >Submit</button>
        <br/>
    </form>
    <hr/>
     <table border="1" align="center" width="80%">
      <tr>
        <th>MemberID</th>
        <th>Firstname</th>
        <th>Lastname</th>
        <th>Age</th>
        <th>Gender</th>
        <th>Delete</th>
      </tr>
      <tr v-for="(member,index) in members" :key="index">
        <td>{{member.member_id}}</td>
        <td>{{member.firstName}}</td>
        <td>{{member.lastName}}</td>
        <td>{{member.age}}</td>
        <td>{{member.Gender.genderName}}</td>
        <td align = "center"><button v-on:click="remove(member.member_id)" >Delete</button></td>
      </tr>
    </table>
  </center>
</template>

<script>
import api from "../api";
export default {
  data() {
    return {
      firstname: "",
      lastname: "",
      age: "",
      gender: "",
      genders: [],
      member: "",
      members: []
    };
  },
  methods: {
    getMember() {
      api.getMember().then(result => {
        console.log(result);
        this.members = result
      });
    },
    getGender() {
      api.getGender().then(result => {
        console.log(result);
        this.genders = result
      });
    },
    submit() {
      api.postMember(
        this.firstname,
        this.lastname,
        this.age,
        this.gender
      );
      window.location.reload()
    },
    remove(id) {
      api.deleteMember(id);
      window.location.reload()
    }
  },
  mounted() {
    this.getMember();
    this.getGender();
  }
}
</script>