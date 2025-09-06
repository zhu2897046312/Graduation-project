import dayjs , { Dayjs } from 'dayjs';
import { ref } from "vue";

export async function formatDate(date : any){
    let newDate = date && ref<Dayjs>(dayjs(date, "YYYY-MM-DD"))
    return newDate.value
  }

export async function getDay(day: any) {
    var today = new Date()
    var targetday_milliseconds = today.getTime() + 1000 * 60 * 60 * 24 * day
    today.setTime(targetday_milliseconds); 
    var tYear : any = today.getFullYear()
    var tMonth : any = today.getMonth()
    var tDate : any = today.getDate()
    tMonth = doHandleMonth(tMonth + 1)
    tDate = doHandleMonth(tDate)
    return tYear + "-" + tMonth + "-" + tDate
}

function doHandleMonth(month: any){
    var m = month;
    if (month.toString().length == 1) {
        m = "0" + month;
    }
    return m;
}