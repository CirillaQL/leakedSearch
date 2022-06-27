import axios from "axios";
export default axios.create({
    baseURL: "http://localhost:33333",
    headers: {
        "Content-type": "application/json"
    }
})