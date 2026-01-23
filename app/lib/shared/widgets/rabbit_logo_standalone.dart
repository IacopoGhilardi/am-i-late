import 'package:flutter/material.dart';
import 'animated_rabbit_logo.dart';

class RabbitLogoStandalone extends StatelessWidget {
  final double size;
  final bool animate;
  final Color? backgroundColor;
  final bool showCircleBackground;

  const RabbitLogoStandalone({
    super.key,
    this.size = 200,
    this.animate = false,
    this.backgroundColor,
    this.showCircleBackground = true,
  });

  @override
  Widget build(BuildContext context) {
    final logo = AnimatedRabbitLogo(size: size * 0.7, animate: animate);

    if (!showCircleBackground) {
      return SizedBox(width: size, height: size, child: Center(child: logo));
    }

    return Container(
      width: size,
      height: size,
      decoration: BoxDecoration(
        color: backgroundColor ?? const Color(0xFF6366F1).withOpacity(0.1),
        shape: BoxShape.circle,
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.1),
            blurRadius: 20,
            offset: const Offset(0, 8),
          ),
        ],
      ),
      child: Container(
        margin: EdgeInsets.all(size * 0.15),
        decoration: const BoxDecoration(
          color: Colors.white,
          shape: BoxShape.circle,
        ),
        child: Center(child: logo),
      ),
    );
  }
}

class LogoShowcasePage extends StatelessWidget {
  const LogoShowcasePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[100],
      appBar: AppBar(
        title: const Text('Am I Late? - Logo Variants'),
        backgroundColor: const Color(0xFF6366F1),
        foregroundColor: Colors.white,
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            _buildSection(
              'Animated Logo (Landing Page)',
              const RabbitLogoStandalone(size: 200, animate: true),
            ),
            const SizedBox(height: 40),
            _buildSection(
              'Static Logo',
              const RabbitLogoStandalone(size: 200, animate: false),
            ),
            const SizedBox(height: 40),
            _buildSection(
              'App Icon Sizes',
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: const [
                  RabbitLogoStandalone(size: 120, animate: false),
                  RabbitLogoStandalone(size: 80, animate: false),
                  RabbitLogoStandalone(size: 60, animate: false),
                ],
              ),
            ),
            const SizedBox(height: 40),
            _buildSection(
              'Logo Only (No Background)',
              Container(
                padding: const EdgeInsets.all(24),
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(16),
                  boxShadow: [
                    BoxShadow(
                      color: Colors.black.withOpacity(0.1),
                      blurRadius: 10,
                      offset: const Offset(0, 4),
                    ),
                  ],
                ),
                child: const RabbitLogoStandalone(
                  size: 150,
                  animate: false,
                  showCircleBackground: false,
                ),
              ),
            ),
            const SizedBox(height: 40),
            _buildSection(
              'On Gradient Background',
              Container(
                padding: const EdgeInsets.all(40),
                decoration: BoxDecoration(
                  gradient: const LinearGradient(
                    colors: [Color(0xFF6366F1), Color(0xFF8B5CF6)],
                    begin: Alignment.topLeft,
                    end: Alignment.bottomRight,
                  ),
                  borderRadius: BorderRadius.circular(16),
                ),
                child: const RabbitLogoStandalone(
                  size: 180,
                  animate: true,
                  backgroundColor: Colors.white,
                ),
              ),
            ),

            const SizedBox(height: 40),

            // Usage instructions
            _buildSection(
              'Usage Notes',
              Container(
                padding: const EdgeInsets.all(20),
                decoration: BoxDecoration(
                  color: Colors.blue[50],
                  borderRadius: BorderRadius.circular(12),
                  border: Border.all(color: Colors.blue[200]!),
                ),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: const [
                    Text(
                      '💡 Tips per l\'uso del logo:',
                      style: TextStyle(
                        fontWeight: FontWeight.bold,
                        fontSize: 16,
                      ),
                    ),
                    SizedBox(height: 12),
                    Text('• Usa la versione animata nella landing page'),
                    Text('• Usa la versione statica per app icon e favicon'),
                    Text('• Il coniglio ha uno smartwatch e fa il gesto OK'),
                    Text('• I colori sono personalizzabili'),
                    Text('• Design completamente originale'),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildSection(String title, Widget content) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          title,
          style: const TextStyle(
            fontSize: 20,
            fontWeight: FontWeight.bold,
            color: Color(0xFF1F2937),
          ),
        ),
        const SizedBox(height: 16),
        Center(child: content),
      ],
    );
  }
}
