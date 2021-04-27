package com.keishi.strava.controller.response;

import com.keishi.strava.domain.MeasurementUnit;
import lombok.Builder;
import lombok.Getter;

@Builder
@Getter
public class WeeklyActivityDetails {

  int distance;

  MeasurementUnit unit;

}
