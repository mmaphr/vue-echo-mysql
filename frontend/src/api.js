import axios from 'axios'

export default{
    getMember() {
        return axios.get("http://localhost:1323/member").then(response => response.data);
    },
    getGender() {
        return axios.get("http://localhost:1323/gender").then(response => response.data);
    },
    postMember(firstname, lastname, age, gender_id) {   
        return axios.post("http://localhost:1323/member",{
            firstname: firstname,
            lastname: lastname,
            age: parseInt(age),
            gender_id: gender_id
        }).then(response => {
            console.log(response.data)
        })
            .catch(function (error) {
            console.log(error)
        });
    },
    deleteMember(id) {
        const url = `http://localhost:1323/member/` + id;
        return axios.delete(url).then(function (response) {
            console.log("Deleted"+response)
        })
            .catch(function (error) {
                console.log(error)
        });
    },
    
}