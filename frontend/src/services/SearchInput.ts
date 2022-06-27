import http from "@/http-common";

/* eslint-disable */
class SearchService {
    getTest(): Promise<any> {
        return http.get("/ping");
    };
    getVideos(value: string): Promise<any> {
        return http.get("/videos/"+value)
    }
}

export default new SearchService();