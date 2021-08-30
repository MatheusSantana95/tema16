package com.matheussantana.devops.tema16.service

import io.gatling.core.Predef._
import io.gatling.core.structure.ScenarioBuilder
import io.gatling.http.Predef._
import io.gatling.http.protocol.HttpProtocolBuilder

import scala.concurrent.duration._
import scala.language.postfixOps

class SampleSimulation extends Simulation {

  val httpConf: HttpProtocolBuilder = http
    .baseUrl("http://localhost:8083")
    .acceptHeader("/calc/history")
    .doNotTrackHeader("1")
    .acceptLanguageHeader("en-US,en;q=0.5")
    .acceptEncodingHeader("gzip, deflate")
    .disableFollowRedirect
    .userAgentHeader("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36")

  val scn: ScenarioBuilder = scenario("Simple Stress Test")
    .exec(http("Http Round Robin 1")
      .get("calc/history")
      .check(status.is(200))
    )

  setUp(
    scn.inject(constantUsersPerSec(100).during(60 seconds))
  ).protocols(httpConf).assertions(global.responseTime.max.lt(1000))

}
