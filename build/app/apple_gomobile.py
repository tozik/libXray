import os
import subprocess
import random
from app.build import Builder

class AppleGoMobileBuilder(Builder):
    def __init__(self, obfuscate=False):
        super().__init__()
        self.obfuscate = obfuscate

    def before_build(self):
        super().before_build()
        self.clean_lib_dirs(["LibXray.xcframework"])
        self.prepare_gomobile()

    def build_obfuscated(self):
        """Специальная сборка с обфускацией"""
        seed = random.randint(100000, 999999)
        cmd = [
            "garble",
            "-seed", str(seed),
            "build",
            "-buildmode=c-archive",
            "-o", "libxray.a",
            "./main.go"
        ]
        result = subprocess.run(cmd, capture_output=True, text=True)
        if result.returncode != 0:
            print(f"Obfuscation failed: {result.stderr}")
            raise Exception("Obfuscated build failed")

    def build(self):
        self.before_build()

        if self.obfuscate:
            self.build_obfuscated()
            
        os.chdir(self.lib_dir)
        cmd = [
            "gomobile",
            "bind",
            "-target", "ios,iossimulator,macos,maccatalyst",
            "-iosversion", "15.0",
            "-o", "LibXray.xcframework"
        ]
        
        if self.obfuscate:
            cmd.extend(["-ldflags", "-s -w"])  # Дополнительная минификация

        result = subprocess.run(cmd, capture_output=True, text=True)
        if result.returncode != 0:
            print(f"Build failed: {result.stderr}")
            raise Exception("Build failed")
