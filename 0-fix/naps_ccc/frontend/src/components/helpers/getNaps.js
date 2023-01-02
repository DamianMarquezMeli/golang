import axios from 'axios';
const baseUrl = process.env.REACT_APP_BASE_URL;

export const getNaps = async () => {
    try {
        const data = await axios.get(`/vianet`,{
            baseURL: baseUrl
        });
        return data.data.data;
    } catch (error) {
        console.log(error);
    }
}
