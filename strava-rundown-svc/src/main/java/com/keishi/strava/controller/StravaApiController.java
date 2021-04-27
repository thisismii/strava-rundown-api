package com.keishi.strava.controller;

import com.keishi.strava.controller.response.WeeklyActivityDetails;
import com.keishi.strava.domain.MeasurementUnit;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class StravaApiController {

  @GetMapping("/activity/details")
  @ResponseBody
  public WeeklyActivityDetails getWeeklyActivityDetails() {
    return WeeklyActivityDetails.builder()
        .distance(10)
        .unit(MeasurementUnit.MILES)
        .build();
  }

}
