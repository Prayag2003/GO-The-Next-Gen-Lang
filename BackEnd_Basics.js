const { resolveSoa } = require("dns");
const express = require("express");  /// import framework
const mongoose = require("mongoose");

const app = express();  // initialise 

const userSchema = mongoose.Schema({
    id: {
        type: Number,
    },
    username: {
        type: String,
    },
    email: {
        type: String,
    },
    phoneno: {
        type: Number,
    },
});

const User = mongoose.model("User", userSchema);


app.use(express.json());  //middleware always run before running of a request

let users = [{
    id: 7,
    username: "jamesbond",
    email: "db09@gmail.com",
    phoneno: 1234567890,
},
{
    id: 7,
    username: "tonystark",
    email: "ly3000@gmail.com",
    phoneno: 1234567899,
}]
// method
app.get("/user", (req, res) => {  // data in req and response in res
    let user = {
        id: 7,
        username: "jamesbond",
        email: "db09@gmail.com",
        phoneno: 1234567890,
    };
    res.send(user);

    app.get("/user/:id", (req, res) => {
        const id = req.params.id;
        let user = users.find(user => user.id == id);
        if (!user) {
            res.status(404).send({ msg: "user doesnt exist" });
        }
        console.log(user);
        res.send(user);
    })
});

app.post("/user", (req, res) => {
    let { id, username, email, phoneno } = req.body;

    let user = {
        id: id,
        username: user,
        email: email,
        phoneno: phoneno,
    };
    users.push(user);
    res.send({ msg: "user created ", data: user });
})
//user id delete -> new data -> add kardo for local host

app.put("user/:id", (req, res) => {
    let id = rew.params.id;
    let { username, email, phoneno } = req.body;
    // remove user
    let user = users.find(user => user.id == id);
    user.splice(users.splice(user), 1);

    // create new user
    let newuser = {
        id: id,
        username: user,
        email: email,
        phoneno: phoneno,
    };
    users.push(newuser);
    res.send({ msg: "user has been updated", data: newuser });

})


app.delete("/user/:id", (req, res) => {
    let id = req.params.id;

    let user = users.find(user => user.id == id);
    user.splice(users.splice(user), 1);

    res.send({ msg: "user deleted", data: user });

})

mongoose.connect('mongodb://localhost:27017/test');

app.listen(3000, () => {    // starting the server
    console.log("Server is Running... ");
});