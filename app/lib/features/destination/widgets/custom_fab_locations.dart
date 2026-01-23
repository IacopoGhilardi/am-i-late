import 'package:flutter/material.dart';

class EndFloatWithTopOffset extends FloatingActionButtonLocation {
  final double topOffset;

  const EndFloatWithTopOffset({this.topOffset = 80});

  @override
  Offset getOffset(ScaffoldPrelayoutGeometry scaffoldGeometry) {
    final double fabX =
        scaffoldGeometry.scaffoldSize.width -
        scaffoldGeometry.floatingActionButtonSize.width -
        16;

    final double fabY = topOffset;

    return Offset(fabX, fabY);
  }
}
