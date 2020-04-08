package cron

import (
    "time"
    "vehicle_system/src/vehicle/cron/cron_lib"
)

const (
    TASK_SPACING_LICENSE = time.Second * 10//一天测试86400
    TASK_EXPIRED = time.Hour * 24 * 365 * 10 //10年
    TASK_SPACING_DEPLOY = 60 //60秒测试
)


func init()  {
    cron := cron_lib.GetTaskScheduler()

    go cron.Start()

    cron.AddFuncSpace(int64(TASK_SPACING_LICENSE), time.Now().UnixNano()+int64(TASK_EXPIRED), LicenseCron)

    //添加指定执行次数任务，并指定每次间隔时间
    //cron.AddFuncSpaceNumber(int64(time.Second*1), 10, func() {
    //    fmt.Println("number 10")
    //})

    //添加一次性任务
    //cron.AddFunc(time.Now().UnixNano()+int64(time.Second*1), func() {
    //    fmt.Println("one second after")
    //})
}
