using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

// Common flags for the command line renderer
// https://knowledge.autodesk.com/support/maya/learn-explore/caas/CloudHelp/cloudhelp/2016/ENU/Maya/files/GUID-0280AB86-8ABE-4F75-B1B9-D5B7DBB7E25A-htm.html

// Render from the command line using mental ray
// https://knowledge.autodesk.com/support/maya/learn-explore/caas/CloudHelp/cloudhelp/2016/ENU/Maya/files/GUID-FCB01DEF-10DF-4222-A047-AF0122782397-htm.html

namespace MayaBatchRender
{
    class RenderSettings
    {
        public int VerbosityLevel { get; set; }
        public string ProjectPath { get; set; }
        public string SceneFolder { get; set; }
        public string SceneName { get; set; }
        public string CameraName { get; set; }
        public string ImageName { get; set; }
        public int PercentResolution { get; set; } // 50 means 50%
    }
    class Renderer
    {
        public Process RenderProcess { get; set; }
        public Renderer(RenderSettings rs)
        {
            RenderProcess = new Process();
            RenderProcess.StartInfo.UseShellExecute = false;
            RenderProcess.StartInfo.RedirectStandardOutput = true;
            RenderProcess.StartInfo.CreateNoWindow = true;
            RenderProcess.StartInfo.FileName = @"C:/Program Files/Autodesk/Maya2016/bin/Render.exe";
            RenderProcess.StartInfo.Arguments += string.Format("-r mr ");
            RenderProcess.StartInfo.Arguments += string.Format("-of png ");
            RenderProcess.StartInfo.Arguments += string.Format("-v {0} ", rs.VerbosityLevel);
            RenderProcess.StartInfo.Arguments += string.Format(@"-proj ""{0}"" ", rs.ProjectPath);
            RenderProcess.StartInfo.Arguments += string.Format(@"-im ""{0}"" ", rs.ImageName);
            RenderProcess.StartInfo.Arguments += string.Format(@"-rd ""{0}"" ", Directory.GetCurrentDirectory() + "/images");
            RenderProcess.StartInfo.Arguments += string.Format("-percentRes {0} ", rs.PercentResolution);
            RenderProcess.StartInfo.Arguments += string.Format("-cam {0} ", rs.CameraName);
            //RenderProcess.StartInfo.Arguments += "-sa 0 ";
            //RenderProcess.StartInfo.Arguments += "-uq " + rs.UnifiedQuality + " ";
            RenderProcess.StartInfo.Arguments += @"""" + rs.ProjectPath + "/" + rs.SceneFolder + "/" + rs.SceneName + ".mb" + @"""";
        }
        public static void ImageMagickGamma(string imageName, float gamma)
        {
            Process process = new Process();

            process.StartInfo.UseShellExecute = false;
            process.StartInfo.RedirectStandardOutput = true;
            process.StartInfo.CreateNoWindow = true;

            process.StartInfo.FileName = @"ImageMagick-6.9.3-0-portable-Q16-x64/convert.exe";
            string imageFullName = Directory.GetCurrentDirectory() + "/images/" + imageName + ".png";
            string nameWithQuotes = @"""" + imageFullName + @"""";
            process.StartInfo.Arguments += nameWithQuotes + " ";
            process.StartInfo.Arguments += string.Format("-gamma {0} ", gamma);
            process.StartInfo.Arguments += nameWithQuotes + " ";

            Console.WriteLine(process.StartInfo.FileName + " " + process.StartInfo.Arguments);

            process.Start();
            process.WaitForExit();
        }

        public static void ImageMagickPNG2JPG(string imageName)
        {
            Process process = new Process();

            process.StartInfo.UseShellExecute = false;
            process.StartInfo.RedirectStandardOutput = true;
            process.StartInfo.CreateNoWindow = true;

            process.StartInfo.FileName = @"ImageMagick-6.9.3-0-portable-Q16-x64/convert.exe";
            string imageFullNameWithoutSuffix = Directory.GetCurrentDirectory() + "/images/" + imageName;
            process.StartInfo.Arguments += @"""" + imageFullNameWithoutSuffix + ".png" + @"""" + " ";
            process.StartInfo.Arguments += string.Format("-quality {0} ", 92);
            process.StartInfo.Arguments += @"""" + imageFullNameWithoutSuffix + ".jpg" + @"""" + " ";

            Console.WriteLine(process.StartInfo.FileName + " " + process.StartInfo.Arguments);

            process.Start();
            process.WaitForExit();
        }
    }
}
