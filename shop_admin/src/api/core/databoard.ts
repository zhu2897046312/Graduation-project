import http from "/@/utils/http.ts";
const base_url = '/missionPlatform/dashboard';


// 关键业务指标
export function getKeyBusinessApi(): Promise<any>{
    return http.post(`${base_url}/getKeyBusiness`)
}

// 功能使用统计
export function getRegUsageApi(data: any){
    return http.post(`${base_url}/getRegUsage`, data)
}

// 功能使用统计
export function getFunctionUsageApi(data: any){
    return http.post(`${base_url}/getFunctionUsage`, data)
}

// 趋势统计
export function getTrendAnalysisApi(data: any){
    return http.post(`${base_url}/getTrendAnalysis`, data)
}

// 小说订单详情
export function getNovelOrderInfoApi(data: any){
    return http.post(`${base_url}/getNovelOrderInfo`, data)
}

// 积分订单详情
export function getCreditOrderInfoApi(data: any){
    return http.post(`${base_url}/getCreditOrderInfo`, data)
}

// ai动画数据详情
export function getVideoProjectInfoApi(data: any){
    return http.post(`${base_url}/getVideoProjectInfo`, data)
}
// 推广详情
export function getPopularizeInfoApi(data: any){
    return http.post(`${base_url}/getPopularizeInfo`, data)
}

// 获取混剪数据详情
export function getMixInfoApi(data: any) {
    return http.post(`${base_url}/getMixInfo`, data)
}

// 获取辅助数据详情
export function getAuxiliaryInfoApi(data: any){
    return http.post(`${base_url}/getAuxiliaryInfo`, data)
}

// ai绘画数据详情
export function getAppAiInfoApi(data: any){
    return http.post(`${base_url}/getAppAiInfo`, data)
}