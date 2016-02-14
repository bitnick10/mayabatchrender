using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;
using YamlDotNet.Serialization;

namespace MayaBatchRender
{
    /// <summary>
    /// Interaction logic for MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public delegate void UpdateTextCallback(string message);
        Thread thread;
        public MainWindow()
        {
            InitializeComponent();
        }

        private void Window_Loaded(object sender, RoutedEventArgs e)
        {
            textBoxConfig.Text = System.IO.File.ReadAllText("config.yaml");

            string yaml = System.IO.File.ReadAllText("config.yaml");
            var deserializer = new Deserializer();
            var input = new StringReader(yaml);
            RenderSettings readyaml = deserializer.Deserialize<RenderSettings>(input);
        }
        private void UpdateText(string message)
        {
            textBoxInfo.AppendText(message);
            textBoxInfo.ScrollToEnd();
        }
        private void Td()
        {
            //RenderSettings rs = new RenderSettings();
            //rs.ProjectPath = @"E:\Pal\art\maya\Songzhou Outskirts";
            //rs.SceneFolder = "scenes/Outskirts";
            //rs.SceneName = "Outskirts";
            //rs.CameraName = "light_AD:camera45gate";
            //rs.PercentResolution = 100;

            string yaml = System.IO.File.ReadAllText("config.yaml");
            var deserializer = new Deserializer();
            var input = new StringReader(yaml);
            RenderSettings rs = deserializer.Deserialize<RenderSettings>(input);

            rs.ImageName = string.Format("{0} {1}", rs.SceneName, DateTime.Now.ToString("yyyy-MM-dd HHmmss"));
            Renderer renderer = new Renderer(rs);
            textBoxInfo.Dispatcher.Invoke(new UpdateTextCallback(UpdateText), renderer.RenderProcess.StartInfo.Arguments);
            renderer.RenderProcess.Start();
            while (!renderer.RenderProcess.StandardOutput.EndOfStream)
            {
                string line = renderer.RenderProcess.StandardOutput.ReadLine();
                textBoxInfo.Dispatcher.Invoke(new UpdateTextCallback(UpdateText), line + Environment.NewLine);
                //Console.WriteLine(line);
            }
            Renderer.ImageMagickGamma(rs.ImageName, 2.2f);
            Renderer.ImageMagickPNG2JPG(rs.ImageName);
            //write log
            textBoxInfo.Dispatcher.BeginInvoke((Action)(() =>
            {
                System.IO.File.WriteAllText("log" + DateTime.Now.ToString("yyyy-MM-dd HHmmss") + ".txt", textBoxInfo.Text);
            }));
            //check shutdown
            CheckOutIsShutdown();
        }

        private void CheckOutIsShutdown()
        {
            checkBoxIsShutdown.Dispatcher.BeginInvoke((Action)(() =>
            {
                if (checkBoxIsShutdown.IsChecked == true)
                {
                    var psi = new ProcessStartInfo("shutdown", "/s /t 0");
                    psi.CreateNoWindow = true;
                    psi.UseShellExecute = false;
                    Process.Start(psi);
                }
            }));
        }

        private void Window_Closing(object sender, System.ComponentModel.CancelEventArgs e)
        {
            if (thread != null)
                thread.Abort();
        }

        private void buttonBeginRender_Click(object sender, RoutedEventArgs e)
        {
            textBoxInfo.Clear();

            thread = new Thread(new ThreadStart(Td));
            thread.Start();
        }

        private void textBoxConfig_TextChanged(object sender, TextChangedEventArgs e)
        {
            System.IO.File.WriteAllText("config.yaml", textBoxConfig.Text);
        }

        private void buttonOpenFolder_Click(object sender, RoutedEventArgs e)
        {
            Process.Start(Directory.GetCurrentDirectory());
        }

        private void buttonInterruptRender_Click(object sender, RoutedEventArgs e)
        {
            if (thread != null)
                thread.Abort();
        }
    }
}
