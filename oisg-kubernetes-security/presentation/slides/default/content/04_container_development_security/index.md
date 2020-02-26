# Container Development Security


# Development Security - DO
- Use a base image
  - Many community or vendor provided images aren't in alignment with best practices
- Relax security to learn, but tighten to deploy
- Use local tools and automation to pre-scan images
- Document and automate security related configurations
- Share & socialize security related learnings

Note:
- have a base image that's approved and hardened by security team makes for easy spot to update 
- Vendors or community may provide a image that just works but maybe you need root to run it, may need to spend some time tweaking the image


# Development Security - DON'T
- Ask for, or expect, security exceptions
- Assume the new technology will “get by” old security policies 
- Create custom images for every new app or build
- Run apps as or containers as root
- Run multiple applications in a container

Note:
- have a base image that's approved and hardened by security team makes for easy spot to update 
- Vendors or community may provide a image that just works but maybe you need root to run it, may need to spend some time tweaking the image