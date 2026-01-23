import 'package:flutter/material.dart';
import 'dart:math' as math;

class AnimatedRabbitLogo extends StatefulWidget {
  final double size;
  final bool animate;

  const AnimatedRabbitLogo({super.key, this.size = 140, this.animate = true});

  @override
  State<AnimatedRabbitLogo> createState() => _AnimatedRabbitLogoState();
}

class _AnimatedRabbitLogoState extends State<AnimatedRabbitLogo>
    with TickerProviderStateMixin {
  late AnimationController _bounceController;
  late AnimationController _rotateController;
  late Animation<double> _bounceAnimation;
  late Animation<double> _rotateAnimation;

  @override
  void initState() {
    super.initState();

    // Bounce animation
    _bounceController = AnimationController(
      duration: const Duration(milliseconds: 1500),
      vsync: this,
    );

    _bounceAnimation = TweenSequence<double>([
      TweenSequenceItem(
        tween: Tween<double>(
          begin: 0,
          end: -10,
        ).chain(CurveTween(curve: Curves.easeOut)),
        weight: 25,
      ),
      TweenSequenceItem(
        tween: Tween<double>(
          begin: -10,
          end: 0,
        ).chain(CurveTween(curve: Curves.bounceOut)),
        weight: 75,
      ),
    ]).animate(_bounceController);

    // Ear rotation animation
    _rotateController = AnimationController(
      duration: const Duration(milliseconds: 2000),
      vsync: this,
    );

    _rotateAnimation = TweenSequence<double>([
      TweenSequenceItem(
        tween: Tween<double>(
          begin: 0,
          end: 0.1,
        ).chain(CurveTween(curve: Curves.easeInOut)),
        weight: 25,
      ),
      TweenSequenceItem(
        tween: Tween<double>(
          begin: 0.1,
          end: -0.1,
        ).chain(CurveTween(curve: Curves.easeInOut)),
        weight: 50,
      ),
      TweenSequenceItem(
        tween: Tween<double>(
          begin: -0.1,
          end: 0,
        ).chain(CurveTween(curve: Curves.easeInOut)),
        weight: 25,
      ),
    ]).animate(_rotateController);

    if (widget.animate) {
      _startAnimations();
    }
  }

  void _startAnimations() {
    Future.delayed(const Duration(milliseconds: 500), () {
      if (mounted) {
        _bounceController.repeat(period: const Duration(seconds: 3));
        _rotateController.repeat(period: const Duration(seconds: 4));
      }
    });
  }

  @override
  void dispose() {
    _bounceController.dispose();
    _rotateController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedBuilder(
      animation: Listenable.merge([_bounceController, _rotateController]),
      builder: (context, child) {
        return Transform.translate(
          offset: Offset(0, _bounceAnimation.value),
          child: CustomPaint(
            size: Size(widget.size, widget.size),
            painter: RabbitLogoPainter(earRotation: _rotateAnimation.value),
          ),
        );
      },
    );
  }
}

class RabbitLogoPainter extends CustomPainter {
  final double earRotation;

  RabbitLogoPainter({this.earRotation = 0});

  @override
  void paint(Canvas canvas, Size size) {
    final center = Offset(size.width / 2, size.height / 2);
    final scale = size.width / 140;

    // Colors
    final rabbitColor = Colors.white;
    final accentColor = const Color(0xFF6366F1); // Primary color
    final innerEarColor = const Color(0xFFFCE7F3); // Light pink
    final shadowColor = Colors.black.withOpacity(0.1);

    // Paint objects
    final rabbitPaint =
        Paint()
          ..color = rabbitColor
          ..style = PaintingStyle.fill;

    final accentPaint =
        Paint()
          ..color = accentColor
          ..style = PaintingStyle.fill;

    final innerEarPaint =
        Paint()
          ..color = innerEarColor
          ..style = PaintingStyle.fill;

    final outlinePaint =
        Paint()
          ..color = accentColor.withOpacity(0.3)
          ..style = PaintingStyle.stroke
          ..strokeWidth = 2 * scale;

    // Shadow
    final shadowPaint =
        Paint()
          ..color = shadowColor
          ..maskFilter = MaskFilter.blur(BlurStyle.normal, 8 * scale);

    canvas.save();
    canvas.translate(center.dx, center.dy);
    canvas.scale(scale);

    // Draw shadow
    canvas.drawCircle(const Offset(0, 35), 25, shadowPaint);

    // Draw left ear
    canvas.save();
    canvas.rotate(earRotation - 0.3);
    _drawEar(
      canvas,
      const Offset(-20, -35),
      rabbitPaint,
      innerEarPaint,
      outlinePaint,
      true,
    );
    canvas.restore();

    // Draw right ear
    canvas.save();
    canvas.rotate(-earRotation + 0.3);
    _drawEar(
      canvas,
      const Offset(20, -35),
      rabbitPaint,
      innerEarPaint,
      outlinePaint,
      false,
    );
    canvas.restore();

    // Draw head (main circle)
    canvas.drawCircle(const Offset(0, 0), 35, rabbitPaint);
    canvas.drawCircle(const Offset(0, 0), 35, outlinePaint);

    // Draw face features
    _drawFace(canvas, accentPaint, rabbitPaint);

    // Draw smartwatch on wrist/arm
    _drawSmartwatch(canvas, accentPaint, rabbitPaint);

    // Draw "OK" hand gesture
    _drawOkGesture(canvas, rabbitPaint, accentPaint, outlinePaint);

    canvas.restore();
  }

  void _drawEar(
    Canvas canvas,
    Offset position,
    Paint mainPaint,
    Paint innerPaint,
    Paint outlinePaint,
    bool isLeft,
  ) {
    final earPath = Path();

    // Outer ear
    earPath.moveTo(position.dx, position.dy);
    earPath.quadraticBezierTo(
      position.dx + (isLeft ? -8 : 8),
      position.dy - 25,
      position.dx,
      position.dy - 35,
    );
    earPath.quadraticBezierTo(
      position.dx + (isLeft ? 8 : -8),
      position.dy - 25,
      position.dx,
      position.dy,
    );

    canvas.drawPath(earPath, mainPaint);
    canvas.drawPath(earPath, outlinePaint);

    // Inner ear
    final innerEarPath = Path();
    innerEarPath.moveTo(position.dx, position.dy - 5);
    innerEarPath.quadraticBezierTo(
      position.dx + (isLeft ? -4 : 4),
      position.dy - 20,
      position.dx,
      position.dy - 28,
    );
    innerEarPath.quadraticBezierTo(
      position.dx + (isLeft ? 4 : -4),
      position.dy - 20,
      position.dx,
      position.dy - 5,
    );

    canvas.drawPath(innerEarPath, innerPaint);
  }

  void _drawFace(Canvas canvas, Paint accentPaint, Paint whitePaint) {
    // Eyes
    final eyePaint =
        Paint()
          ..color = const Color(0xFF1F2937)
          ..style = PaintingStyle.fill;

    // Left eye
    canvas.drawCircle(const Offset(-12, -5), 4, eyePaint);
    // Right eye
    canvas.drawCircle(const Offset(12, -5), 4, eyePaint);

    // Eye highlights
    canvas.drawCircle(const Offset(-11, -6), 1.5, whitePaint);
    canvas.drawCircle(const Offset(13, -6), 1.5, whitePaint);

    // Nose
    final nosePath = Path();
    nosePath.moveTo(0, 2);
    nosePath.lineTo(-3, 8);
    nosePath.lineTo(3, 8);
    nosePath.close();
    canvas.drawPath(nosePath, accentPaint);

    // Smile
    final smilePath = Path();
    smilePath.moveTo(-8, 12);
    smilePath.quadraticBezierTo(0, 16, 8, 12);

    final smilePaint =
        Paint()
          ..color = const Color(0xFF1F2937)
          ..style = PaintingStyle.stroke
          ..strokeWidth = 2
          ..strokeCap = StrokeCap.round;

    canvas.drawPath(smilePath, smilePaint);

    // Whiskers
    final whiskerPaint =
        Paint()
          ..color = const Color(0xFF9CA3AF)
          ..style = PaintingStyle.stroke
          ..strokeWidth = 1.5
          ..strokeCap = StrokeCap.round;

    // Left whiskers
    canvas.drawLine(const Offset(-15, 0), const Offset(-28, -2), whiskerPaint);
    canvas.drawLine(const Offset(-15, 4), const Offset(-28, 4), whiskerPaint);

    // Right whiskers
    canvas.drawLine(const Offset(15, 0), const Offset(28, -2), whiskerPaint);
    canvas.drawLine(const Offset(15, 4), const Offset(28, 4), whiskerPaint);
  }

  void _drawSmartwatch(Canvas canvas, Paint accentPaint, Paint whitePaint) {
    // Smartwatch on left wrist (bottom left area)
    final watchCenter = const Offset(-25, 25);

    // Watch band
    final bandPaint =
        Paint()
          ..color = accentPaint.color
          ..style = PaintingStyle.stroke
          ..strokeWidth = 4;

    canvas.drawLine(
      Offset(watchCenter.dx - 2, watchCenter.dy - 6),
      Offset(watchCenter.dx - 2, watchCenter.dy - 15),
      bandPaint,
    );

    // Watch face
    canvas.drawRRect(
      RRect.fromRectAndRadius(
        Rect.fromCenter(center: watchCenter, width: 10, height: 10),
        const Radius.circular(2),
      ),
      accentPaint,
    );

    // Watch screen highlight
    final screenPaint =
        Paint()
          ..color = const Color(0xFF818CF8)
          ..style = PaintingStyle.fill;

    canvas.drawRRect(
      RRect.fromRectAndRadius(
        Rect.fromCenter(center: watchCenter, width: 7, height: 7),
        const Radius.circular(1),
      ),
      screenPaint,
    );

    // Time display on watch (small clock icon)
    final clockPaint =
        Paint()
          ..color = whitePaint.color
          ..style = PaintingStyle.stroke
          ..strokeWidth = 0.8;

    canvas.drawCircle(watchCenter, 2, clockPaint);
    canvas.drawLine(
      watchCenter,
      Offset(watchCenter.dx, watchCenter.dy - 1.5),
      clockPaint,
    );
    canvas.drawLine(
      watchCenter,
      Offset(watchCenter.dx + 1, watchCenter.dy),
      clockPaint,
    );
  }

  void _drawOkGesture(
    Canvas canvas,
    Paint mainPaint,
    Paint accentPaint,
    Paint outlinePaint,
  ) {
    // Arm/hand position (bottom right)
    final handCenter = const Offset(28, 20);

    // Small arm
    final armPaint =
        Paint()
          ..color = mainPaint.color
          ..style = PaintingStyle.stroke
          ..strokeWidth = 8
          ..strokeCap = StrokeCap.round;

    canvas.drawLine(
      const Offset(25, 10),
      Offset(handCenter.dx - 5, handCenter.dy),
      armPaint,
    );

    // OK gesture circle (thumb and index finger)
    canvas.drawCircle(handCenter, 6, mainPaint);
    canvas.drawCircle(handCenter, 6, outlinePaint);

    // Inner circle (the "O" part)
    final okCirclePaint =
        Paint()
          ..color = accentPaint.color
          ..style = PaintingStyle.stroke
          ..strokeWidth = 2;

    canvas.drawCircle(handCenter, 3, okCirclePaint);

    // Three other fingers pointing up
    final fingerPaint =
        Paint()
          ..color = mainPaint.color
          ..style = PaintingStyle.stroke
          ..strokeWidth = 3
          ..strokeCap = StrokeCap.round;

    canvas.drawLine(
      Offset(handCenter.dx - 3, handCenter.dy - 5),
      Offset(handCenter.dx - 4, handCenter.dy - 12),
      fingerPaint,
    );
    canvas.drawLine(
      Offset(handCenter.dx, handCenter.dy - 6),
      Offset(handCenter.dx, handCenter.dy - 14),
      fingerPaint,
    );
    canvas.drawLine(
      Offset(handCenter.dx + 3, handCenter.dy - 5),
      Offset(handCenter.dx + 4, handCenter.dy - 12),
      fingerPaint,
    );
  }

  @override
  bool shouldRepaint(RabbitLogoPainter oldDelegate) {
    return oldDelegate.earRotation != earRotation;
  }
}
